package discordha

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

type HA interface {
	// AddHandler makes a discordgo handler HA, it will only be ran on one replica
	AddHandler(handler interface{}) func()
	// CacheRead allows to read a cached object from etcd shared across all replicas
	CacheRead(cache, key string, want interface{}) (interface{}, error)
	// CacheWrite allows to read a cached object from etcd shared across all replicas
	CacheWrite(cache, key string, data interface{}, ttl time.Duration) error
	// Stop should be run when the program terminates
	Stop()

	// The following functions are due to change

	// LockVoice locks a voice channel for a specific module and Guild, returns true if successful, this function may change soon!
	LockVoice(moduleID, guildID string) (bool, error)
	// LockVoice locks a voice channel for a specific module and Guild, this function may change soon!
	UnlockVoice(moduleID, guildID string) error
	// SendVoiceCommand sends a string command to the instance handling the voice channel
	// These can be received using WatchVoiceCommands
	// this function may change soon!
	SendVoiceCommand(command VoiceCommand) error
	// WatchVoiceCommands gives a channel with commands transmitted by SendVoiceCommand
	// this function may change soon!
	WatchVoiceCommands(ctx context.Context, moduleID string) chan VoiceCommand
}

// HA is a helper struct for high available discordgo using etcd
type HAInstance struct {
	config      *Config
	etcd        *clientv3.Client
	locksMutex  sync.Mutex
	locks       map[string]clientv3.LeaseID
	bgContext   context.Context
	name        string
	concurrency *concurrency.Session

	// internal isLeader, if true it should be confirmed by etcd before acting like one!
	isLeader bool
}

// Config contains the configuration for HA
type Config struct {
	Session                          *discordgo.Session
	HA                               bool
	LockUpdateInterval               time.Duration
	LockTTL                          time.Duration
	EtcdEndpoints                    []string
	Context                          context.Context
	Log                              log.Logger
	VerboseLevel                     int
	DoNotParticipateInLeaderElection bool
}

// New gives a HA instance for a given configuration
func New(c *Config) (HA, error) {
	if !c.HA {
		return &HAInstance{
			config: c,
		}, nil
	}

	if c.Session == nil {
		return nil, errors.New("no discordgo session passed")
	}

	// set the defaults
	if c.LockUpdateInterval == 0 {
		c.LockUpdateInterval = time.Second * 1
	}
	if c.LockTTL == 0 {
		c.LockTTL = time.Second * 10
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.EtcdEndpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	concur, err := concurrency.NewSession(client, concurrency.WithTTL(int(c.LockUpdateInterval.Seconds())))
	var s = &HAInstance{
		config:      c,
		etcd:        client,
		locks:       map[string]clientv3.LeaseID{},
		bgContext:   c.Context,
		name:        fmt.Sprintf("%d", rand.Intn(9999999)),
		concurrency: concur,
	}

	// start DiscordHA leader election
	if !c.DoNotParticipateInLeaderElection {
		go func() {
			for {
				err := s.ElectLeader(c.Context)
				if err == nil {
					break // became the leader, end this Go routine
				}
				s.config.Log.Println("Error in etcd leader election", err)
				time.Sleep(time.Second)
			}
		}()
	}

	//update the locks so they do not die on long lived command runs
	go s.lockUpdateLoop(c.Context)
	if s.config.VerboseLevel >= 1 {
		go s.logLoop(c.Context, time.Hour) // log number of locks every hour
	}

	return s, nil
}

// Stop should be run when the program terminates
func (h *HAInstance) Stop() {
	err := h.ResignLeader(context.TODO())
	if err != nil {
		h.config.Log.Println("Error resigning leader", err)
	}
}
