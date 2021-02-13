package types

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramChannel int64
type DiscordChannel string

type GoferConfig struct {
	Channels         map[string]ChannelMapping `json:"channels"`
	DiscordApiToken  string                             `json:"discordApiToken"`
	TelegramApiToken string                             `json:"telegramApiToken"`
}


type Context struct {
	Discord *discordgo.Session
	Telegram *tgbotapi.BotAPI
	Config GoferConfig
	Listener map[string]Listeners
}

type Listeners struct {
	File ListenerFileCb
	Message ListenerMessageCb
	Misc ListenerMiscCb
	Photo ListenerPhotoCb
	EditMessage ListenerEditMessageCb
}


type ChannelMapping struct {
	DiscordChanId string
	TelegramChanId int64
	MatrixChanId string
}
