package main

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramChannel int64
type DiscordChannel string

type GoferConfig struct {
	Channels         map[TelegramChannel]DiscordChannel `json:"channels"`
	DiscordApiToken  string                             `json:"discordApiToken"`
	TelegramApiToken string                             `json:"telegramApiToken"`
}


type Context struct {
	discordBot *discordgo.Session
	telegramBot *tgbotapi.BotAPI
}