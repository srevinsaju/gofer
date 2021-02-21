package matrix

import (
	"errors"
	"github.com/srevinsaju/gofer/types"
)

func GetChannels(chanId string, channels map[string]types.ChannelMapping) (types.ChannelMapping, error) {
	for _, v := range channels {
		if v.MatrixChanId == chanId {
			return v, nil
		}
	}
	return types.ChannelMapping{}, errors.New("not found")
}