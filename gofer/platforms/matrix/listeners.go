package matrix

import (
	"fmt"

	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/format"
	"maunium.net/go/mautrix/id"

	"github.com/srevinsaju/gofer/types"
)

func SendMessage(ctx types.Context, channel types.ChannelMapping, message types.GoferMessage) error {
	var strMessage string
	if message.ReplyTo != "" {
		strMessage = fmt.Sprintf(
			"<blockquote><i>In reply to <b>%s</b></i>: %s</blockquote>\n<b>%s</b>: %s",
			message.ReplyTo, message.ReplyToMessage, message.From, message.Message)
	} else {
		strMessage = fmt.Sprintf("<b>%s</b>: %s", message.From, message.Message)
	}

	content := format.RenderMarkdown(strMessage, true, true)
	_, err := ctx.Matrix.SendMessageEvent(id.RoomID(channel.MatrixChanId), event.EventMessage, &content)
	if err != nil {
		logger.Warnf("Failed to send message to matrix server")
		return err
	}
	return nil
}

func SendEdit(ctx types.Context, channel types.ChannelMapping, message types.GoferEditedMessage) error {
	strMessage := fmt.Sprintf("<b>%s</b>: * [edit] <i>%s</i>", message.From, message.Message)

	content := format.RenderMarkdown(strMessage, true, true)
	_, err := ctx.Matrix.SendMessageEvent(id.RoomID(channel.MatrixChanId), event.EventMessage, &content)
	if err != nil {
		logger.Warnf("Failed to send message to matrix server")
		return err
	}
	return nil
}

func SendPhoto(ctx types.Context, channel types.ChannelMapping, photo types.GoferPhoto) error {
	strMessage := ""
	if photo.ReplyToMessage != "" {
		strMessage = fmt.Sprintf("* <i>In reply to %s</i>\n<blockquote>%s</blockquote>\n\n<i>%s sent a photo</i>", photo.ReplyTo, photo.ReplyToMessage, photo.From)
	} else {
		strMessage = fmt.Sprintf("* <i>%s sent a photo</i> %s", photo.From, photo.Message)
	}

	content := format.RenderMarkdown(strMessage, true, true)
	link, err := ctx.Matrix.UploadLink(photo.Url)
	if err != nil {
		logger.Warnf("Couldn't upload photo to matrix homeserver")
		return err
	}

	_, err = ctx.Matrix.SendImage(id.RoomID(channel.MatrixChanId), "", link.ContentURI)
	if err != nil {
		logger.Warnf("Failed to send image to matrix")
		return err
	}

	_, err = ctx.Matrix.SendMessageEvent(id.RoomID(channel.MatrixChanId), event.EventMessage, &content)
	if err != nil {
		logger.Warnf("Failed to send message to matrix server")
		return err
	}
	return nil
}

func SendFile(ctx types.Context, channel types.ChannelMapping, file types.GoferFile) error {
	strMessage := fmt.Sprintf("* <i><b>%s</b> sent a file on %s</i>", file.From, file.Origin)
	if file.Url != "" {
		strMessage = strMessage + "\n\n" + file.Url
	}
	if file.Message != "" {
		strMessage = strMessage + "\n\n" + file.Message
	}

	content := format.RenderMarkdown(strMessage, true, true)
	link, err := ctx.Matrix.UploadLink(file.Url)
	if err != nil {
		logger.Warnf("Couldn't upload photo to matrix homeserver")
		return err
	}

	_, err = ctx.Matrix.SendImage(id.RoomID(channel.MatrixChanId), "", link.ContentURI)
	if err != nil {
		logger.Warnf("Failed to send image to matrix")
		return err
	}

	_, err = ctx.Matrix.SendMessageEvent(id.RoomID(channel.MatrixChanId), event.EventMessage, &content)
	if err != nil {
		logger.Warnf("Failed to send message to matrix server")
		return err
	}
	return nil
}

func SendMisc(ctx types.Context, channel types.ChannelMapping, misc types.GoferMisc) error {
	strMessage := fmt.Sprintf("* <i><b>%s</b> sent a %s %s on %s</i>", misc.From, misc.Identifier, misc.Message, misc.Origin)

	content := format.RenderMarkdown(strMessage, true, true)
	_, err := ctx.Matrix.SendMessageEvent(id.RoomID(channel.MatrixChanId), event.EventMessage, &content)
	if err != nil {
		logger.Warnf("Failed to send message to matrix server")
		return err
	}
	return nil
}
