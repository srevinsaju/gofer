package main

import "errors"

func reverseGet(arr map[TelegramChannel]DiscordChannel, value DiscordChannel) (TelegramChannel, error) {
	for k, v := range arr {
		if v == value {
			return k, nil
		}
	}
	return TelegramChannel(0), errors.New("didnt find the channel")
}
