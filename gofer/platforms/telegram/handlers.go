package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/srevinsaju/gofer/orchestra"
	"github.com/srevinsaju/gofer/types"
)

func EventHandler(ctx types.Context) {
	logger.Infof("Authorized on account %s", ctx.Telegram.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := ctx.Telegram.GetUpdatesChan(u)
	if err != nil {
		logger.Fatal("Failed to get updates channel")
		return
	}

	for update := range updates {
		if update.EditedMessage == nil && update.Message == nil { // ignore any non-Message Updates
			continue
		}
		
		if update.EditedMessage != nil {
			// this contains an edited message
			channel, err := GetChannels(update.EditedMessage.Chat.ID, ctx.Config.Channels)
			if err != nil {
				continue
			}

			orchestra.SendEditMessageTo(ctx, channel, "telegram", types.GoferEditedMessage{
				From:           update.EditedMessage.From.FirstName,
				Message:        update.EditedMessage.Text,
				Origin:         "telegram",
			})
			continue
		}

		if update.Message != nil {
			channel, err := GetChannels(update.Message.Chat.ID, ctx.Config.Channels)
			if err != nil {
				continue
			}

			replyMessage := ""
			replyTo := ""
			if update.Message.ReplyToMessage != nil {
				replyMessage = update.Message.ReplyToMessage.Text
				replyTo = update.Message.ReplyToMessage.From.FirstName
			}

			// check if pictures / files
			if update.Message.Photo != nil {
				photos := *update.Message.Photo
				var bestPhoto tgbotapi.PhotoSize
				width := 0

				// get the biggest photo
				for i := range photos {
					if photos[i].Width > width {
						bestPhoto = photos[i]
					}
				}

				url, err := ctx.Telegram.GetFileDirectURL(bestPhoto.FileID)
				if err != nil {
					logger.Warnf("Couldnt get direct URL, %s", err)
					continue
				}
				orchestra.SendPhotoTo(ctx, channel, "telegram", types.GoferPhoto{
					From:           update.Message.From.FirstName,
					Url:            url,
					Message:        update.Message.Caption,
					ReplyTo:        replyTo,
					Origin:         "telegram",
					ReplyToMessage: replyMessage,
				})
				continue
			}

			if update.Message.Document != nil {
				url, err := ctx.Telegram.GetFileDirectURL(update.Message.Document.FileID)
				if err != nil {
					logger.Warnf("Couldnt get direct URL, %s", err)
					continue
				}
				orchestra.SendFileTo(ctx, channel, "telegram", types.GoferFile{
					From:           update.Message.From.FirstName,
					Url:            url,
					Message:        update.Message.Text,
					ReplyTo:        replyTo,
					Origin:         "telegram",
					ReplyToMessage: replyMessage,
				})
				continue
			}

			if update.Message.Sticker != nil {
				orchestra.SendMiscTo(ctx, channel, "telegram", types.GoferMisc{
					From:           update.Message.From.FirstName,
					Url:            "",
					Message:        update.Message.Sticker.Emoji,
					ReplyTo:        replyTo,
					Origin:         "telegram",
					Identifier:     "sticker",
					ReplyToMessage: replyMessage,
				})

			}

			if update.Message.Text == "" {
				continue
			}

			logger.Infof("telegram:[%s] %s", update.Message.From.FirstName, update.Message.Text)
			orchestra.SendMessageTo(ctx, channel, "telegram", types.GoferMessage{
				From:           update.Message.From.FirstName,
				Message:        update.Message.Text,
				ReplyTo:        replyTo,
				ReplyToMessage: replyMessage,
				Origin:         "telegram",
			})
		}
	}
}