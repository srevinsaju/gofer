package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func CreateConfig() {
	cfg := GoferConfig{}
	cfg.Channels = map[TelegramChannel]DiscordChannel{}

	var inputBuf string

	// get Telegram Api Token from @BotFather
	fmt.Print("Enter Telegram API Token: ")
	_, err := fmt.Scanln(&inputBuf)
	if err != nil {
		logger.Fatal(err)
		return
	}
	fmt.Println("")
	cfg.TelegramApiToken = inputBuf

	// get DiscordApiToken from Discord Application portal
	fmt.Print("Enter Discord API Token: ")
	_, err = fmt.Scanln(&inputBuf)
	if err != nil {
		logger.Fatal(err)
		return
	}
	fmt.Println("")
	cfg.DiscordApiToken = inputBuf

	for true {
		fmt.Println("Enter Telegram ChanID followed by Discord Channel Id separated by")
		fmt.Println("a comma, for example: -432232xx,12345667833453633553530")
		_, err = fmt.Scanln(&inputBuf)

		if inputBuf == "EXIT" || err != nil {
			break
		}

		mapping := strings.Split(inputBuf, ",")

		// do some checks on the telegram Channel ID
		telegramChanId, err := strconv.Atoi(mapping[0])
		if err != nil {
			logger.Warnf("%s is not a valid telegram channel id", mapping[0])
			continue
		}

		telegramChanIdTyped := TelegramChannel(telegramChanId)
		discordChanIdTyped := DiscordChannel(mapping[1])
		cfg.Channels[telegramChanIdTyped] = discordChanIdTyped

	}

	outputBytes, err := json.Marshal(cfg)
	err = ioutil.WriteFile("gofer.json", outputBytes, 0644)
	if err != nil {
		logger.Fatal(err)
		return
	}

}

func ConfigFromFile(filepath string) (GoferConfig, error) {
	rawData, err := ioutil.ReadFile(filepath)
	var cfg GoferConfig
	err = json.Unmarshal(rawData, &cfg)
	if err != nil {
		logger.Fatal(err)
		return GoferConfig{}, err
	}
	return cfg, nil
}
