package discordha

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

// etcd states to store in value
const (
	statusNone     = "0"
	statusHandling = "1"
	statusOk       = "2"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (h *HAInstance) ElectLeader(ctx context.Context) error {
	if !h.config.HA {
		// Non HA, development instance probably
		return nil
	}
	e := concurrency.NewElection(h.concurrency, "/discordha-election/")
	if err := e.Campaign(ctx, h.name); err != nil {
		return err
	}

	h.isLeader = true

	return nil
}

func (h *HAInstance) ResignLeader(ctx context.Context) error {
	if !h.config.HA {
		// Non HA, development instance probably
		return nil
	}
	e := concurrency.NewElection(h.concurrency, "/discordha-election/")
	if err := e.Resign(ctx); err != nil {
		return err
	}

	h.isLeader = true

	return nil
}

func (h *HAInstance) AmLeader(ctx context.Context) bool {
	if !h.config.HA {
		// Non HA, development instance probably
		return true
	}
	if !h.isLeader {
		// saves one etcd roundtrip
		return false
	}
	e := concurrency.NewElection(h.concurrency, "/discordha-election/")
	resp, err := e.Leader(ctx)
	if err != nil {
		h.config.Log.Println("AmLeader error", err)
		return false
	}
	return bytes.Equal(resp.Kvs[0].Value, []byte(h.name))
}

func (h *HAInstance) lockUpdateLoop() {
	for {
		time.Sleep(h.config.LockUpdateInterval)
		h.locksMutex.Lock()
		for _, lease := range h.locks {
			err := h.keepAlive(lease)
			if err != nil {
				h.config.Log.Printf("Etcd keepalive error: %q\n", err)
			}
		}
		h.locksMutex.Unlock()
	}
}

func (h *HAInstance) logLoop() {
	c := 0
	for {
		time.Sleep(time.Minute)
		h.locksMutex.Lock()
		c++
		if c > 100 {
			h.config.Log.Printf("I own %d locks\n", len(h.locks))
			c = 0
		}
		h.locksMutex.Unlock()
	}
}

// Lock tries to acquire a lock on an event, it will return true if
// the instance that requests it may process the request.
func (h *HAInstance) Lock(obj interface{}) (bool, error) {
	if !h.config.HA {
		// Non HA, development instance probably
		return true, nil
	}

	hash, err := h.getObjectHash(obj)
	if err != nil {
		h.config.Log.Printf("Hash error:%q\n", err)
		return false, err
	}
	key := fmt.Sprintf("/locks/%s", hash)
	return h.lockKey(key, true)
}

func (h *HAInstance) lockKey(key string, waitForFailure bool) (bool, error) {
	grant, err := h.etcd.Grant(h.bgContext, int64(h.config.LockTTL.Seconds()))
	if err != nil {
		return false, err
	}

	txn, err := h.etcd.Txn(h.bgContext).
		// txn value comparisons are lexical
		If(clientv3.Compare(clientv3.Value(key), ">", statusNone)).
		Else(clientv3.OpPut(key, statusHandling, clientv3.WithLease(grant.ID))).
		Commit()

	if err != nil {
		return false, err
	}

	// if clientv3.Compare(clientv3.Value(key), ">", statusNone) is true
	if txn.Succeeded {
		// Lock exists!
		if !waitForFailure {
			return false, nil
		}
		ctx, cancel := context.WithCancel(h.bgContext)
		defer cancel()

		w := h.etcd.Watch(ctx, key)
		for wresp := range w {
			if wresp.Canceled {
				return h.lockKey(key, waitForFailure) // attempt watch again!
			}
			for _, ev := range wresp.Events {
				if string(ev.Kv.Value) == statusOk {
					// other server succeeded!
					return false, nil
				}
				if ev.Type == clientv3.EventTypeDelete {
					return h.lockKey(key, waitForFailure) // re-lock!
				}
			}
		}
		return false, nil
	}

	h.locksMutex.Lock()
	h.locks[key] = grant.ID
	h.locksMutex.Unlock()

	return true, nil
}

// Unlock will release a lock on an event
func (h *HAInstance) Unlock(obj interface{}) error {
	if !h.config.HA {
		// Non HA, development instance probably
		return nil
	}

	hash, err := h.getObjectHash(obj)
	if err != nil {
		h.config.Log.Printf("Hash error:%q\n", err)
		return err
	}
	key := fmt.Sprintf("/locks/%s", hash)

	return h.unlockKey(key)
}

func (h *HAInstance) unlockKey(key string) error {
	h.keepAlive(h.locks[key]) // keep lock in etcd till it expires so all servers catch up
	_, err := h.etcd.Put(h.bgContext, key, statusOk, clientv3.WithLease(h.locks[key]))
	if err != nil {
		h.config.Log.Printf("Failed to set status OK: %q retrying", err)
		time.Sleep(5 * time.Second)
		return h.unlockKey(key)
	}

	h.locksMutex.Lock()
	delete(h.locks, key)
	h.locksMutex.Unlock()

	return nil
}

func (h *HAInstance) keepAlive(leaseID clientv3.LeaseID) error {
	_, err := h.etcd.KeepAlive(h.bgContext, leaseID)
	return err
}

func (h *HAInstance) getObjectHash(v interface{}) (string, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	hasher := sha256.New()
	hasher.Write(jsonData)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil
}
