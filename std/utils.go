package std

import (
	"strings"

	"github.com/andersfylling/disgord"
)

// CopyMsgEvt allows you to copy a message event such that you can edit the content using middlewares.
// if you don't create a copy, the same change will be present in _every_ other handler that depends on the
// same event type.
//
// If you're making a bot, and you only care about the data after it has been processed, then there should
// be no need to create a copy. Just remember that you're dealing with pointer, and a change in one handler/middleware
// changes the data for every other handler/middleware.
func CopyMsgEvt(evt interface{}) interface{} {
	switch t := evt.(type) {
	case *disgord.MessageCreate:
		return &disgord.MessageCreate{
			Message: disgord.DeepCopy(t.Message).(*disgord.Message),
			ShardID: t.ShardID,
		}
	case *disgord.MessageUpdate:
		return &disgord.MessageUpdate{
			Message: disgord.DeepCopy(t.Message).(*disgord.Message),
			ShardID: t.ShardID,
		}
	case *disgord.MessageDelete:
		return &disgord.MessageDelete{
			MessageID: t.MessageID,
			ChannelID: t.ChannelID,
			GuildID:   t.GuildID,
			ShardID:   t.ShardID,
		}
	case *disgord.MessageDeleteBulk:
		return &disgord.MessageDeleteBulk{
			MessageIDs: t.MessageIDs,
			ChannelID:  t.ChannelID,
			ShardID:    t.ShardID,
		}
	case *disgord.MessageReactionAdd:
		return &disgord.MessageReactionAdd{
			PartialEmoji: disgord.DeepCopy(t.PartialEmoji).(*disgord.PartialEmoji),
			MessageID:    t.MessageID,
			ChannelID:    t.ChannelID,
			UserID:       t.UserID,
			ShardID:      t.ShardID,
		}
	case *disgord.MessageReactionRemove:
		return &disgord.MessageReactionRemove{
			PartialEmoji: disgord.DeepCopy(t.PartialEmoji).(*disgord.PartialEmoji),
			MessageID:    t.MessageID,
			ChannelID:    t.ChannelID,
			UserID:       t.UserID,
			ShardID:      t.ShardID,
		}
	case *disgord.MessageReactionRemoveAll:
		return &disgord.MessageReactionRemoveAll{
			MessageID: t.MessageID,
			ChannelID: t.ChannelID,
			ShardID:   t.ShardID,
		}
	case *disgord.MessageReactionRemoveEmoji:
		return &disgord.MessageReactionRemoveEmoji{
			MessageID: t.MessageID,
			GuildID:   t.GuildID,
			ChannelID: t.ChannelID,
			Emoji:     disgord.DeepCopy(t.Emoji).(*disgord.Emoji),
			ShardID:   t.ShardID,
		}
	}

	// TODO: logging might be useful
	return nil
}

func mentionString(id disgord.Snowflake) string {
	return "<@" + id.String() + ">"
}

func getMsg(evt interface{}) (msg *disgord.Message) {
	switch t := evt.(type) {
	case *disgord.MessageCreate:
		msg = t.Message
	case *disgord.MessageUpdate:
		msg = t.Message
	default:
		msg = nil
	}

	return msg
}

func messageHasPrefix(evt interface{}, prefix string) interface{} {
	var msg *disgord.Message
	if msg = getMsg(evt); msg == nil {
		return nil
	}

	content := strings.TrimSpace(msg.Content)
	if !strings.HasPrefix(content, prefix) {
		return nil
	}

	return evt
}

func messageIsBot(evt interface{}, isBot bool) interface{} {
	var msg *disgord.Message
	if msg = getMsg(evt); msg == nil {
		return nil
	}

	if msg.Author != nil && msg.Author.Bot != isBot {
		return nil
	}

	return evt
}

func messageIsWebhook(evt interface{}, isWebhook bool) interface{} {
	var msg *disgord.Message
	if msg = getMsg(evt); msg == nil {
		return nil
	}

	// Webhook ID is only set if message is sourced from a webhook.
	// https://discord.com/developers/docs/resources/channel#message-object-message-structure
	// "You can tell if a message is generated by a webhook by checking for the webhook_id
	// on the message object."
	if msg.WebhookID.IsZero() == isWebhook {
		return nil
	}

	return evt
}
