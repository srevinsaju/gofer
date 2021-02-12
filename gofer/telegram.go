package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TelegramMediaWrapper(
	discordBot *discordgo.Session,
	telegramBot *tgbotapi.BotAPI,
	fileId string,
	discordChanId string,
) {
	doc, err := telegramBot.GetFileDirectURL(fileId)
	if err != nil {
		logger.Warnf("[TelegramBot] Failed to get file from telegram, %s", err)
		return
	}
	image := discordgo.MessageEmbedImage{URL: doc}
	_, err = discordBot.ChannelMessageSendEmbed(discordChanId, &discordgo.MessageEmbed{Image: &image})
	if err != nil {
		logger.Warnf("[DiscordBot] Failed to send file from telegram to %s, %s", discordChanId, err)
		return
	}
}

func TelegramOnMessageHandler(telegramBot *tgbotapi.BotAPI, update tgbotapi.Update, discordBot *discordgo.Session, config GoferConfig) {
	var chanId string
	if val, ok := config.Channels[TelegramChannel(update.Message.Chat.ID)]; ok {
		chanId = string(val)
	} else {
		logger.Infof("[TelegramBot] Received an event from unrecognized channel")
		return
	}

	if update.Message.Photo != nil {
		logger.Infof("[TelegramBot] [%s] %s", update.Message.From.FirstName, "Sending Photo")
		photoPointer := *update.Message.Photo
		for photoIdx := range photoPointer {
			photoFile := photoPointer[photoIdx]
			TelegramMediaWrapper(discordBot, telegramBot, photoFile.FileID, chanId)
		}
	} else if update.Message.Document != nil {
		logger.Infof("[TelegramBot] [%s] %s", update.Message.From.FirstName, "Sending File")
		TelegramMediaWrapper(discordBot, telegramBot, update.Message.Document.FileID, chanId)
	}

	if update.Message.Text == "" {
		return
	}

	logger.Infof("[TelegramBot] [%s] %s", update.Message.From.FirstName, update.Message.Text)
	_, err := discordBot.ChannelMessageSend(
		chanId,
		fmt.Sprintf("**%s**: %s", update.Message.From.FirstName, update.Message.Text))
	if err != nil {
		logger.Infof("[DiscordBot] Failed to send message %s", err)
	}
}

func TelegramOnEditedMessageHandler(_ *tgbotapi.BotAPI, update tgbotapi.Update, discordBot *discordgo.Session, config GoferConfig) {
	var chanId string
	if val, ok := config.Channels[TelegramChannel(update.EditedMessage.Chat.ID)]; ok {
		chanId = string(val)
	} else {
		logger.Infof("[TelegramBot] Received an event from unrecognized channel")
		return
	}
	logger.Infof("[TelegramBot] [%s] (edited) %s", update.EditedMessage.From.FirstName, update.EditedMessage.Text)

	_, err := discordBot.ChannelMessageSend(
		chanId,
		fmt.Sprintf("**%s**: * %s", update.EditedMessage.From.FirstName, update.EditedMessage.Text))
	if err != nil {
		logger.Infof("[DiscordBot] Failed to send message %s", err)
	}
}

func TelegramEventHandler(telegramBot *tgbotapi.BotAPI, discordBot *discordgo.Session, config GoferConfig) {

	logger.Infof("[TelegramBot] Authorized on account %s", telegramBot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := telegramBot.GetUpdatesChan(u)
	if err != nil {
		logger.Fatal("Failed to get updates channel")
		return
	}

	for update := range updates {
		if update.EditedMessage == nil && update.Message == nil { // ignore any non-Message Updates
			continue
		}
		var handler func(telegramBot *tgbotapi.BotAPI, update tgbotapi.Update, discordBot *discordgo.Session, config GoferConfig)
		if update.EditedMessage == nil {
			// its a message event
			handler = TelegramOnMessageHandler
		} else if update.Message == nil {
			handler = TelegramOnEditedMessageHandler
		} else {
			continue
		}

		handler(telegramBot, update, discordBot, config)

	}
}
