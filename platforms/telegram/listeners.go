package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/srevinsaju/gofer/types"
)


func SendMessage(ctx types.Context, channel types.ChannelMapping, message types.GoferMessage ) error {
	var strMessage string
	if message.ReplyTo != "" {
		strMessage = fmt.Sprintf(
			"<i>In reply to <b>%s</b></i>: %s\n<b>%s</b>: %s",
			message.ReplyTo, message.ReplyToMessage, message.From, message.Message)
	} else {
		strMessage = fmt.Sprintf("<b>%s</b>: %s", message.From, message.Message)
	}
	msg := tgbotapi.NewMessage(channel.TelegramChanId, strMessage)
	msg.ParseMode = "html"
	_, err := ctx.Telegram.Send(msg)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}

func SendEdit(ctx types.Context, channel types.ChannelMapping, message types.GoferEditedMessage ) error {
	strMessage := fmt.Sprintf("<b>%s</b>: * [edit] <i>%s</i>", message.From, message.Message)

	msg := tgbotapi.NewMessage(channel.TelegramChanId, strMessage)
	msg.ParseMode = "html"
	_, err := ctx.Telegram.Send(msg)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}

func SendPhoto(ctx types.Context, channel types.ChannelMapping, photo types.GoferPhoto ) error {
	strMessage := ""
	if photo.ReplyToMessage != "" {
		strMessage = fmt.Sprintf("* <i>In reply to %s</i>\n%s\n\n<i>%s sent a photo</i>\n<a href=\"%s\">.</a>", photo.ReplyTo, photo.ReplyToMessage, photo.From, photo.Url)
	} else {
		strMessage = fmt.Sprintf("* <i>%s sent a photo</i>\n<a href=\"%s\">.</a> %s", photo.From, photo.Url, photo.Message)
	}

	msg := tgbotapi.NewMessage(channel.TelegramChanId, strMessage)
	msg.ParseMode = "html"
	_, err := ctx.Telegram.Send(msg)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}

func SendFile(ctx types.Context, channel types.ChannelMapping, file types.GoferFile ) error {
	strMessage := fmt.Sprintf("* <i><b>%s</b> sent a file on %s</i>", file.From, file.Origin)
	if file.Url != "" {
		strMessage = strMessage + "\n\n" + file.Url
	}
	if file.Message != "" {
		strMessage = strMessage + "\n\n" + file.Message
	}

	msg := tgbotapi.NewMessage(channel.TelegramChanId, strMessage)
	msg.ParseMode = "html"
	_, err := ctx.Telegram.Send(msg)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}

func SendMisc(ctx types.Context, channel types.ChannelMapping, misc types.GoferMisc ) error {
	strMessage := fmt.Sprintf("* <i><b>%s</b> sent a %s %s on %s</i>", misc.From, misc.Identifier, misc.Message, misc.Origin)

	msg := tgbotapi.NewMessage(channel.TelegramChanId, strMessage)
	msg.ParseMode = "html"
	_, err := ctx.Telegram.Send(msg)
	if err != nil {
		logger.Infof("Failed to send message %s", err)
		return err
	}
	return nil
}