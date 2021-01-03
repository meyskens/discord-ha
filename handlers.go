package discordha

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *HAInstance) AddHandler(handler interface{}) func() {
	var wrappedHandler interface{}
	switch v := handler.(type) {
	case func(*discordgo.Session, *discordgo.ChannelCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelCreate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.ChannelDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelDelete) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.ChannelPinsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.ChannelUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Connect):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Connect) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Disconnect):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Disconnect) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Event):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Event) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildBanAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildBanAdd) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildBanRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildBanRemove) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildCreate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildDelete) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildEmojisUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildIntegrationsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMemberAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMemberRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMemberUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMembersChunk):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildRoleCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildRoleDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildRoleUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageAck):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageAck) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageCreate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageDelete) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageDeleteBulk):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageDeleteBulk) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageReactionAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
			// reactions are non timestamped and should only be handled by one instance
			if !h.AmLeader(context.TODO()) {
				return
			}
			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageReactionRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
			// reactions are non timestamped and should only be handled by one instance
			if !h.AmLeader(context.TODO()) {
				return
			}
			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageReactionRemoveAll):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageReactionRemoveAll) {
			// reactions are non timestamped and should only be handled by one instance
			if !h.AmLeader(context.TODO()) {
				return
			}
			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.PresenceUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.PresenceUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.PresencesReplace):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.PresencesReplace) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.RateLimit):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.RateLimit) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Ready):
		wrappedHandler = v // this should be sent to all instances
	case func(*discordgo.Session, *discordgo.RelationshipAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.RelationshipAdd) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.RelationshipRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.RelationshipRemove) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Resumed):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Resumed) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.TypingStart):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.TypingStart) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserGuildSettingsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserNoteUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserNoteUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserSettingsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.VoiceServerUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.VoiceStateUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.WebhooksUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.WebhooksUpdate) {
			if ok, err := h.Lock(e); !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(e)

			v(s, e)
		}
	}

	if wrappedHandler == nil {
		panic("Provided handler not recognised") // TODO: handle this better
	}

	return h.config.Session.AddHandler(wrappedHandler)
}
