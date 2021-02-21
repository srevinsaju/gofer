package matrix

import (
	"github.com/srevinsaju/gofer/orchestra"
	"github.com/srevinsaju/gofer/types"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"strings"
	"time"
)

func EventHandler (ctx types.Context) {

	syncer := ctx.Matrix.Syncer.(*mautrix.DefaultSyncer)


	startTime := time.Now().UnixNano() / 1_000_000

	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
		if evt.Timestamp < startTime {
			// Ignore events from before the program started
			return
		}


		mapping, err := GetChannels(evt.RoomID.String(), ctx.Config.Channels)
		if err != nil {
			return
		}

		nick := strings.Split(evt.Sender.String(), ":")[0]
		nick = strings.TrimPrefix(nick, "@")

		if evt.Content.AsMember().Displayname != "" {
			nick = evt.Content.AsMember().Displayname
		}

		if evt.Sender == ctx.Matrix.UserID {
			return
		}

		logger.Infof("%s: %s", evt.Content.AsMessage().MsgType, evt.Content.AsMessage().Body)


		switch evt.Content.AsMessage().MsgType {
		case "m.text":
			logger.Infof("matrix:[%s] %s", nick, evt.Content.AsMessage().Body)
			orchestra.SendMessageTo(ctx, mapping, "matrix", types.GoferMessage{
				From:           nick,
				Message:        evt.Content.AsMessage().Body,
				ReplyTo:        evt.Content.AsMessage().To.String(),
				ReplyToMessage: evt.Content.AsMessage().GetReplyTo().String(),
				Origin:         "matrix",
			})


		case "m.image" :

			mxcUrl, err := evt.Content.AsMessage().URL.Parse()
			if err != nil {
				logger.Warnf("Couldn't get photo URL, %s ", err)
				return
			}
			photoUrl := ctx.Matrix.GetDownloadURL(mxcUrl)

			orchestra.SendPhotoTo(ctx, mapping, "matrix", types.GoferPhoto{
				From:           nick,
				Url:            photoUrl,
				Message:        evt.Content.AsMessage().Body,
				ReplyTo:        evt.Content.AsMessage().To.String(),
				Origin:         "matrix",
				ReplyToMessage: evt.Content.AsMessage().GetReplyTo().String(),
			})


		case "m.file":
			mxcUrl, err := evt.Content.AsMessage().URL.Parse()
			if err != nil {
				logger.Warnf("Couldn't get file URL, %s ", err)
				return
			}
			fileUrl := ctx.Matrix.GetDownloadURL(mxcUrl)
			orchestra.SendFileTo(ctx, mapping, "matrix", types.GoferFile{
				From:           nick,
				Url:            fileUrl,
				Message:        evt.Content.AsMessage().Body,
				ReplyTo:        evt.Content.AsMessage().To.String(),
				Origin:         "matrix",
				ReplyToMessage: evt.Content.AsMessage().GetReplyTo().String(),
			})

		}

	})

	err := ctx.Matrix.Sync()

	if err != nil {
		panic(err)
	}
}