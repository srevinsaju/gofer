package matrix

import (
	"fmt"
	"github.com/srevinsaju/gofer/orchestra"
	"github.com/srevinsaju/gofer/types"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
)

func EventHandler (ctx types.Context) {

	syncer := ctx.Matrix.Syncer.(*mautrix.DefaultSyncer)

	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {

		mapping, err := GetChannels(evt.RoomID.String(), ctx.Config.Channels)
		if err != nil {
			return
		}

		nick := evt.Content.AsMember().Displayname


		switch evt.Type.String() {
		case "m.text":
			logger.Infof("matrix:[%s] %s", nick, evt.Content.AsMessage().Body)
			orchestra.SendMessageTo(ctx, mapping, "matrix", types.GoferMessage{
				From:           nick,
				Message:        evt.Content.AsMessage().Body,
				ReplyTo:        evt.Content.AsMessage().To.String(),
				ReplyToMessage: evt.Content.AsMessage().GetReplyTo().String(),
				Origin:         "matrix",
			})


		case "m.photo" :
			photoUrl, err  := evt.Content.AsMessage().GetFile().URL.Parse()
			if err != nil {
				logger.Warnf("Couldn't get photo URL")
				return
			}
			orchestra.SendPhotoTo(ctx, mapping, "matrix", types.GoferPhoto{
				From:           nick,
				Url:            photoUrl.String(),
				Message:        evt.Content.AsMessage().Body,
				ReplyTo:        evt.Content.AsMessage().To.String(),
				Origin:         "matrix",
				ReplyToMessage: evt.Content.AsMessage().GetReplyTo().String(),
			})


		case "m.file":
			fileUrl, err := evt.Content.AsMessage().GetFile().URL.Parse()
			if err != nil {
				logger.Warnf("Couldn't get file URL")
				return
			}
			orchestra.SendFileTo(ctx, mapping, "matrix", types.GoferFile{
				From:           nick,
				Url:            fileUrl.String(),
				Message:        evt.Content.AsMessage().Body,
				ReplyTo:        evt.Content.AsMessage().To.String(),
				Origin:         "matrix",
				ReplyToMessage: evt.Content.AsMessage().GetReplyTo().String(),
			})

		}
		fmt.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)
	})

	err := ctx.Matrix.Sync()

	if err != nil {
		panic(err)
	}
}