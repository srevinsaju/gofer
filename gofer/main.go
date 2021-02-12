package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/withmandala/go-log"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.New(os.Stdout)

func main() {
	// get last command
	command := os.Args[len(os.Args)-1]
	if command == "create" {
		CreateConfig()
		logger.Info("Config file written. Call `gofer` like")
		logger.Infof("~ $ gofer path/to/config.json")
		return
	}

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
	// create the telegram bot
	telegramBot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		logger.Fatal(err)
	}

	// create the discord bot
	discordBotToken := cfg.DiscordApiToken
	if discordBotToken == "" {
		logger.Fatal("config.discordApiToken is not provided")
	}
	discordBot, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		logger.Fatal("[DiscordBot] Failed to initialize discord bot")
		return
	}

	telegramBot.Debug = false
	discordBot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		DiscordOnMessageHandler(s, m, telegramBot, cfg)
	})

	discordBot.Identify.Intents = discordgo.IntentsGuildMessages

	logger.Infof("[DiscordBot] Bot is now running.  Press CTRL-C to exit.")
	go TelegramEventHandler(telegramBot, discordBot, cfg)
	// Open a websocket connection to Discord and begin listening.
	err = discordBot.Open()
	if err != nil {
		logger.Infof("error opening connection, %s", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = discordBot.Close()
	if err != nil {
		logger.Fatal(err)
	}
}
