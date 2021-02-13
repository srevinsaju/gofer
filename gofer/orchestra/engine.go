package orchestra

import "github.com/srevinsaju/gofer/types"


func SendMessageTo(ctx types.Context, mapping types.ChannelMapping, except string, message types.GoferMessage ) {
	for k, v := range ctx.Listener {
		if k == except {
			continue
		}

		go v.Message(ctx, mapping, message)
	}
}

func SendEditMessageTo(ctx types.Context, mapping types.ChannelMapping, except string, message types.GoferEditedMessage ) {
	for k, v := range ctx.Listener {
		if k == except {
			continue
		}

		go v.EditMessage(ctx, mapping, message)
	}
}



func SendFileTo(ctx types.Context, mapping types.ChannelMapping, except string, file types.GoferFile ) {
	for k, v := range ctx.Listener {
		if k == except {
			continue
		}

		go v.File(ctx, mapping, file)
	}
}

func SendPhotoTo(ctx types.Context, mapping types.ChannelMapping, except string, photo types.GoferPhoto ) {
	for k, v := range ctx.Listener {
		if k == except {
			continue
		}

		go v.Photo(ctx, mapping, photo)
	}
}

func SendMiscTo(ctx types.Context, mapping types.ChannelMapping, except string, misc types.GoferMisc ) {
	for k, v := range ctx.Listener {
		if k == except {
			continue
		}

		go v.Misc(ctx, mapping, misc)
	}
}