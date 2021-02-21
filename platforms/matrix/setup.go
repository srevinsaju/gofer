package matrix

import (
	"github.com/srevinsaju/gofer/types"
	"github.com/withmandala/go-log"
	"maunium.net/go/mautrix"
	"os"
)

var logger = log.New(os.Stdout)

func Setup (ctx *types.Context) {
	homeserver := ctx.Config.MatrixHomeServer
	password := ctx.Config.MatrixPassword
	username := ctx.Config.MatrixUsername

	if username == "" && password == "" && homeserver == "" {
		// the user doesnt want to bridge to matrix network
		return
	}

	if username == "" || password == "" || homeserver == "" {
		// something information is not provided
		logger.Fatal("Invalid params for Matrix bot. Couldn't find either username, password or homeserver")
		return
	}

	client, err := mautrix.NewClient(homeserver, "", "")
	if err != nil {
		logger.Fatalf("Couldn't connect to matrix %s, %s", homeserver, err)
		return
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type:             "m.login.password",
		Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: username},
		Password:         password,
		StoreCredentials: true,
	})
	if err != nil {
		logger.Fatalf("Couldn't login to Matrix %s with %s, %s", homeserver, username, err)
		return
	}

	logger.Infof("Authorization successful as %s on %s", username, homeserver)



	ctx.Matrix = client

}