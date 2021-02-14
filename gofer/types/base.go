package types

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramChannel int64
type DiscordChannel string

type GoferConfig struct {
	Channels         map[string]ChannelMapping 			`json:"channels"`
	DiscordApiToken  string                             `json:"discord_api_token,omitempty"`
	TelegramApiToken string                             `json:"telegram_api_token,omitempty"`
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
	DiscordChanId string `json:"discord_chan_id,omitempty"`
	TelegramChanId int64 `json:"telegram_chan_id,omitempty"`
	MatrixChanId string `json:"matrix_chan_id,omitempty"`
}
