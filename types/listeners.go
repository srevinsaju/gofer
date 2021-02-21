package types

type ListenerMessageCb func(ctx Context, channel ChannelMapping, message GoferMessage) error
type ListenerEditMessageCb func(ctx Context, channel ChannelMapping, message GoferEditedMessage) error
type ListenerPhotoCb func(ctx Context, channel ChannelMapping, photo GoferPhoto ) error
type ListenerFileCb func(ctx Context, channel ChannelMapping, file GoferFile) error
type ListenerMiscCb func(ctx Context, channel ChannelMapping, misc GoferMisc) error