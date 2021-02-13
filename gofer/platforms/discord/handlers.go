package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/srevinsaju/gofer/orchestra"
	"github.com/srevinsaju/gofer/types"
)

// EventHandler function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func EventHandler(ctx *types.Context, s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// do not process Direct messages
	if m.Type != 0 {
		return
	}

	channel, err := GetChannels(m.ChannelID, ctx.Config.Channels)
	if err != nil {
		return
	}

	if m.Message.Attachments != nil {
		for idx := range m.Message.Attachments {

			url := m.Message.Attachments[idx].URL
			orchestra.SendFileTo(*ctx, channel, "discord", types.GoferFile{
				From:           m.Message.Member.Nick,
				Url:            url,
				Message:        m.Message.ContentWithMentionsReplaced(),
				ReplyTo:        "",  // fixme
				Origin:         "discord",
				ReplyToMessage: "",
			})
		}
		return
	}

	if m.Message.Content == "" {
		return
	}

	logger.Infof("[%s] %s", m.Author.Username, m.Message.Content)
	orchestra.SendMessageTo(*ctx, channel, "discord", types.GoferMessage{
		From:           m.Message.Member.Nick,
		Message:        m.Message.ContentWithMentionsReplaced(),
		ReplyTo:        "", // fixme
		ReplyToMessage: "",
		Origin:         "discord",
	})
}

