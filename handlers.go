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
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.ChannelDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelDelete) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.ChannelPinsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.ChannelUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.ChannelUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Connect):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Connect) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Disconnect):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Disconnect) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Event):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Event) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildBanAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildBanAdd) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildBanRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildBanRemove) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildCreate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildDelete) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildEmojisUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildIntegrationsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMemberAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
			// GuildMemberAdd are non timestamped and should only be handled by one instance
			if !h.AmLeader(context.TODO()) {
				return
			}
			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMemberRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMemberUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildMembersChunk):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildRoleCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildRoleDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildRoleUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.GuildUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.GuildUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageAck):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageAck) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageCreate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageDelete):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageDelete) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.MessageDeleteBulk):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.MessageDeleteBulk) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

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
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.PresenceUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.PresenceUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.PresencesReplace):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.PresencesReplace) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.RateLimit):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.RateLimit) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Ready):
		wrappedHandler = v // this should be sent to all instances
	case func(*discordgo.Session, *discordgo.RelationshipAdd):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.RelationshipAdd) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.RelationshipRemove):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.RelationshipRemove) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.Resumed):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.Resumed) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.TypingStart):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.TypingStart) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserGuildSettingsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserNoteUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserNoteUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserSettingsUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.UserUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.UserUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.VoiceServerUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.VoiceStateUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.WebhooksUpdate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.WebhooksUpdate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	case func(*discordgo.Session, *discordgo.InteractionCreate):
		wrappedHandler = func(s *discordgo.Session, e *discordgo.InteractionCreate) {
			ok, key, err := h.Lock(e)
			if !ok {
				if err != nil {
					log.Println(err)
				}
				return
			}
			defer h.Unlock(key)

			v(s, e)
		}
	}

	if wrappedHandler == nil {
		panic("Provided handler not recognised") // TODO: handle this better
	}

	return h.config.Session.AddHandler(wrappedHandler)
}
