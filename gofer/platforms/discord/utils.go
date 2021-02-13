package discord

import (
	"errors"
	"github.com/srevinsaju/gofer/types"
	"github.com/withmandala/go-log"
	"os"
)

var logger = log.New(os.Stdout)


func GetChannels(chanId string, channels map[string]types.ChannelMapping) (types.ChannelMapping, error) {
	for _, v := range channels {
		if v.DiscordChanId == chanId {
			return v, nil
		}
	}
	return types.ChannelMapping{}, errors.New("not found")
}