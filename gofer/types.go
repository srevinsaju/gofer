package main

type TelegramChannel int64
type DiscordChannel string

type GoferConfig struct {
	Channels 			map[TelegramChannel]DiscordChannel	`json:"channels"`
	DiscordApiToken		string								`json:"discordApiToken"`
	TelegramApiToken	string								`json:"telegramApiToken"`
}