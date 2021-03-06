package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/srevinsaju/gofer/types"
	"mime"
	"net/http"
)


type DiscordWebHook struct {
	Content string `json:"content"`
	Username string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}


func SendMessage(ctx types.Context, channel types.ChannelMapping, message types.GoferMessage ) error {
	if channel.DiscordWebhook != "" {
		return SendWebHookMessage(ctx, channel, message)
	}

	var strMessage string
	if message.ReplyTo != "" {
		strMessage = fmt.Sprintf("> _In reply to **%s**: %s_\n**%s**: %s",
			message.ReplyTo, message.ReplyToMessage, message.From, message.Message)
	} else {
		strMessage = fmt.Sprintf("**%s**: %s", message.From, message.Message)
	}
	_, err := ctx.Discord.ChannelMessageSend(channel.DiscordChanId, strMessage)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}


func SendWebHookMessage(ctx types.Context, channel types.ChannelMapping, message types.GoferMessage ) error {
	var strMessage string
	if message.ReplyTo != "" {
		strMessage = fmt.Sprintf("> _In reply to **%s**: %s_\n%s",
			message.ReplyTo, message.ReplyToMessage, message.Message)
	} else {
		strMessage = message.Message
	}
	webhook := DiscordWebHook{
		Content:   strMessage,
		Username:  message.From,
		AvatarUrl: message.FromUserProfilePic,
	}
	byteData, err := json.Marshal(webhook)
	if err != nil {
		logger.Warn("Failed to marshal json webhook data, ", webhook)
		return err
	}
	resp, err := http.Post(channel.DiscordWebhook, "application/json", bytes.NewReader(byteData))
	if err != nil {
		logger.Warnf("Failed to send POST request to webhook, %s", webhook)
		return err
	}
	logger.Infof("Discord webhook returned %s", resp.StatusCode)

	return nil
}

func SendImage(ctx types.Context, channel types.ChannelMapping, photo types.GoferPhoto ) error {
	description := ""
	if photo.ReplyTo != "" {
		description = fmt.Sprintf("In reply to _%s_ \n> _%s_ \n\n %s", photo.ReplyTo, photo.ReplyToMessage, photo.Message )
	} else {
		description = fmt.Sprintf("%s", photo.Message)
	}

	r, err := http.Get(photo.Url)
	if err != nil {
		logger.Warnf("Failed to get %s", photo.Url)
		return err
	}
	if r.StatusCode != 200 {
		return nil
	}

	logger.Debugf("Content type: %s", r.Header.Get("Content-Type"))

	contentType := r.Header.Get("Content-Type")
	extension, err := mime.ExtensionsByType(contentType)

	_, err = ctx.Discord.ChannelMessageSendComplex(channel.DiscordChanId, &discordgo.MessageSend{
		Content: fmt.Sprintf("_**%s** sent a photo_\n %s", photo.From, description),
		Files: []*discordgo.File{
			{
				Name:        fmt.Sprintf("image.%s", extension[0]),
				ContentType: contentType,
				Reader:      r.Body,
			},
		},
	})

	if err != nil {
		logger.Warnf("Failed to send image to %s, %s", channel.DiscordChanId, err)
		return err
	}
	return nil
}

func SendFile(ctx types.Context, channel types.ChannelMapping, file types.GoferFile ) error {
	strMessage := fmt.Sprintf("_* %s sent a file on %s_", file.From, file.Origin)

	if file.Url != "" {
		strMessage = strMessage + "\n\n" + file.Url
	}
	if file.Message != "" {
		strMessage = strMessage + "\n\n" + file.Message
	}

	_, err := ctx.Discord.ChannelMessageSend(channel.DiscordChanId, strMessage)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}


func SendMisc(ctx types.Context, channel types.ChannelMapping, misc types.GoferMisc ) error {
	strMessage := fmt.Sprintf("_* %s sent a %s %s on %s_", misc.From, misc.Identifier, misc.Message, misc.Origin)
	_, err := ctx.Discord.ChannelMessageSend(channel.DiscordChanId, strMessage)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}

func SendEdit(ctx types.Context, channel types.ChannelMapping, message types.GoferEditedMessage) error {
	strMessage := fmt.Sprintf("**%s**: %s", message.From, message.Message)
	_, err := ctx.Discord.ChannelMessageSend(channel.DiscordChanId, strMessage)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}