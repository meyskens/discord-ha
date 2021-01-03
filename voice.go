package discordha

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/coreos/etcd/clientv3"
)

type VoiceCommand struct {
	ChannelID string `json:"channelID"`
	File      string `json:"file"`
	UserID    string `json:"userID"`
}

// LockVoice locks a voice channel ID, returns true if successful
func (h *HAInstance) LockVoice(channelID string) (bool, error) {
	return h.lockKey(fmt.Sprintf("voice-%s", channelID), false)
}

// UnlockVoice unlocks a voice channel ID
func (h *HAInstance) UnlockVoice(channelID string) error {
	return h.unlockKey(fmt.Sprintf("voice-%s", channelID))
}

// SendVoiceCommand sends a string command to the instance handling the voice channel
// These can be received using WatchVoiceCommands
func (h *HAInstance) SendVoiceCommand(channelID string, command VoiceCommand) error {
	grant, err := h.etcd.Grant(context.TODO(), int64(30))
	if err != nil {
		return err
	}
	cmd, _ := json.Marshal(command)
	_, err = h.etcd.Put(context.TODO(), fmt.Sprintf("/voice/command/%s/%d", channelID, rand.Intn(9999999)), string(cmd), clientv3.WithLease(grant.ID))
	return err
}

// WatchVoiceCommands gives a channel with commands transmitted by SendVoiceCommand
func (h *HAInstance) WatchVoiceCommands(ctx context.Context, channelID string) chan VoiceCommand {
	out := make(chan VoiceCommand)
	w := h.etcd.Watch(ctx, fmt.Sprintf("/voice/command/%s/", channelID), clientv3.WithPrefix())
	go func() {
		for wresp := range w {
			if wresp.Canceled {
				close(out)
				break
			}
			for _, ev := range wresp.Events {
				if ev.IsCreate() {
					cmd := VoiceCommand{}
					err := json.Unmarshal(ev.Kv.Value, &cmd)
					if err == nil {
						out <- cmd
					}
				}
			}
		}
	}()

	return out
}
