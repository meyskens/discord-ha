package discordha

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// ErrorCacheKeyNotExist is the error the cache returns if a key does not exist
var ErrorCacheKeyNotExist = errors.New("Cache key does not exist")

// CacheRead reads a key from a specific cache, returns ErrorCacheKeyNotExist if not found
func (h *HAInstance) CacheRead(cache, key string, want interface{}) (interface{}, error) {
	if !h.config.HA {
		return nil, ErrorCacheKeyNotExist
	}
	resp, err := h.etcd.Get(context.TODO(), fmt.Sprintf("/cache/%s/%s", cache, key))
	if err != nil {
		return nil, err
	}

	if resp.Count < 1 {
		return nil, ErrorCacheKeyNotExist
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &want)
	if err != nil {
		return nil, err
	}

	return want, nil
}

// CacheWrite writes an object to a specific cache with a specific key, will be purged after TTL expires
func (h *HAInstance) CacheWrite(cache, key string, data interface{}, ttl time.Duration) error {
	if !h.config.HA {
		return nil
	}
	grant, err := h.etcd.Grant(context.TODO(), int64(ttl.Seconds()))
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = h.etcd.Put(context.TODO(), fmt.Sprintf("/cache/%s/%s", cache, key), string(jsonData), clientv3.WithLease(grant.ID))
	return err
}
