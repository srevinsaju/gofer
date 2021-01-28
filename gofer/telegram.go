package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


func TelegramOnMessageHandler(telegramBot *tgbotapi.BotAPI, discordBot *discordgo.Session, config GoferConfig) {

	logger.Infof("[TelegramBot] Authorized on account %s", telegramBot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := telegramBot.GetUpdatesChan(u)
	if err != nil {
		logger.Fatal("Failed to get updates channel")
		return
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		var chanId string
		if val, ok := config.Channels[TelegramChannel(update.Message.Chat.ID)]; ok {
			chanId = string(val)
		} else {
			logger.Infof("[TelegramBot] Received an event from unrecognized channel")
			continue
		}
		logger.Infof("[TelegramBot] [%s] %s", update.Message.From.FirstName, update.Message.Text)

		_, err = discordBot.ChannelMessageSend(
			chanId,
			fmt.Sprintf("**%s**: %s", update.Message.From.FirstName ,update.Message.Text))
		if err != nil {
			logger.Infof("[DiscordBot] Failed to send message %s", err)
		}

	}
}

