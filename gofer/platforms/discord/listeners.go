package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/srevinsaju/gofer/types"
)


func SendMessage(ctx types.Context, channel types.ChannelMapping, message types.GoferMessage ) error {
	var strMessage string
	if message.ReplyTo != "" {
		strMessage = fmt.Sprintf("> %s \n**%s**: %s", message.ReplyTo, message.From, message.Message)
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


func SendImage(ctx types.Context, channel types.ChannelMapping, photo types.GoferPhoto ) error {
	image := discordgo.MessageEmbedImage{URL: photo.Url}

	description := ""
	if photo.ReplyTo != "" {
		description = fmt.Sprintf("In reply to _%s_ \n> _%s_ \n\n %s", photo.ReplyTo, photo.ReplyToMessage, photo.Message )
	} else {
		description = fmt.Sprintf("%s", photo.Message)
	}
	embed := discordgo.MessageEmbed{
		Type:        "",
		Title:       photo.From,
		Description: description,
		Timestamp:   "",
		Footer:      nil,
		Image:       &image,
	}
	_, err := ctx.Discord.ChannelMessageSendEmbed(
		channel.DiscordChanId,
		&embed,
	)
	if err != nil {
		logger.Warnf("Failed to send file to %s, %s", channel.DiscordChanId, err)
		return err
	}
	return nil
}

func SendFile(ctx types.Context, channel types.ChannelMapping, file types.GoferFile ) error {
	strMessage := fmt.Sprintf("_* %s sent a file on %s_", file.From, file.Origin)
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