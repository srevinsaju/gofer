package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/srevinsaju/gofer/platforms/discord"
	"github.com/srevinsaju/gofer/platforms/telegram"
	"github.com/srevinsaju/gofer/types"
	"github.com/withmandala/go-log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = log.New(os.Stdout)

func main() {
	// get last command
	command := os.Args[len(os.Args)-1]

	if command == "gofer" {
		// the user has not provided any commands along with the executable name
		// so, we should show the usage
		logger.Info("gofer : yet another telegram - discord bridge")
		logger.Info("")
		logger.Info("To load an existing configuration: ")
		logger.Info("  $ gofer path/to/config.json")
		logger.Info("")
		logger.Info("To create a new configuration in current directory:")
		logger.Info("  $ gofer create")
		return

	}

	if _, err := os.Stat(command); os.IsNotExist(err) {
		logger.Fatal("The specified path does not exist")
	}

	goferCfgFile := command

	cfg, err := ConfigFromFile(goferCfgFile)
	if err != nil {
		logger.Fatal(err)
	}

	telegramBotToken := cfg.TelegramApiToken

	// create the discord bot
	discordListeners := types.Listeners{
		File:        discord.SendFile,
		Message:     discord.SendMessage,
		Misc:        discord.SendMisc,
		Photo:       discord.SendImage,
		EditMessage: discord.SendEdit,
	}

	telegramListeners := types.Listeners{
		File:        telegram.SendFile,
		Message:     telegram.SendMessage,
		Misc:        telegram.SendMisc,
		Photo:       telegram.SendPhoto,
		EditMessage: telegram.SendEdit,
	}

	ctx := &types.Context{
		Config: cfg,
		Listener: map[string]types.Listeners{
			"discord":  discordListeners,
			"telegram": telegramListeners,
		},
	}

	if ctx.Config.DiscordApiToken != "" {

		discordBotToken := cfg.DiscordApiToken
		if discordBotToken == "" {
			logger.Fatal("config.discordApiToken is not provided")
		}
		discordBot, err := discordgo.New("Bot " + discordBotToken)
		if err != nil {
			logger.Fatal("[DiscordBot] Failed to initialize discord bot")
			return
		}
		ctx.Discord = discordBot

		ctx.Discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			discord.EventHandler(ctx, s, m)
		})
		ctx.Discord.Identify.Intents = discordgo.IntentsGuildMessages

		logger.Infof("Authorized on Discord Account")
	}

	if ctx.Config.TelegramApiToken != "" {
		// create the telegram bot
		telegramBot, err := tgbotapi.NewBotAPI(telegramBotToken)
		if err != nil {
			logger.Fatal(err)
		}
		telegramBot.Debug = false
		ctx.Telegram = telegramBot
	}

	logger.Info("Starting Telegram event handler")
	go telegram.EventHandler(*ctx)
	logger.Info("Starting Discord event handler")
	go ctx.Discord.Open()

	for true {
		time.Sleep(time.Second * 2)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = ctx.Discord.Close()
	if err != nil {
		logger.Fatal(err)
	}
}
