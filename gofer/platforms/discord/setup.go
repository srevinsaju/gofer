package discord

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)


func DiscordSetup(ctx *Context) error {

	discordBotToken := ctx.config.DiscordApiToken
	if discordBotToken == "" {
		err := errors.New("[DiscordBot] config[\"discord_api_token\"] is not provided")
		return err
	}

	discordBot, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		return err
	}

	ctx.discordBot = discordBot

	ctx.discordBot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		DiscordOnMessageHandler(ctx, s, m)
	})
	ctx.discordBot.Identify.Intents = discordgo.IntentsGuildMessages

	return nil
}


