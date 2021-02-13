package telegram

import (
	"errors"
	"github.com/srevinsaju/gofer/types"
)

func GetChannels(chanId int64, channels map[string]types.ChannelMapping) (types.ChannelMapping, error) {
	for _, v := range channels {
		if v.TelegramChanId == chanId {
			return v, nil
		}
	}
	return types.ChannelMapping{}, errors.New("not found")
}