package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// DiscordOnMessageHandler function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func DiscordOnMessageHandler(
	s *discordgo.Session, m *discordgo.MessageCreate, telegramBot *tgbotapi.BotAPI, config GoferConfig) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// do not process Dm messages
	if m.Type != 0 {
		return
	}

	key, err := reverseGet(config.Channels, DiscordChannel(m.Message.ChannelID))
	if err != nil {
		logger.Info("[DiscordBot] Received an event from unregistered channel")
		return
	}

	msg := tgbotapi.NewMessage(
		int64(key),
		fmt.Sprintf("*%s*: %s", m.Author.Username,  m.ContentWithMentionsReplaced()),
	)
	// set markdown mode for formatting the username
	msg.ParseMode = "markdown"

	_, err = telegramBot.Send(msg)
	if err != nil {
		logger.Warnf("[TelegramBot] Failed to send message from discord, %s", err)
	}
}