package types

type GoferMessage struct {
	From string
	FromUserProfilePic string
	Message string
	ReplyTo string
	ReplyToMessage string
	Origin string
}

type GoferPhoto struct {
	From string
	Url string
	Message string
	ReplyTo string
	Origin string
	ReplyToMessage string
}

type GoferFile struct {
	From string
	Url string
	Message string
	ReplyTo string
	Origin string
	ReplyToMessage string
}

type GoferMisc struct {
	From string
	Url string
	Message string
	ReplyTo string
	Origin string
	Identifier string
	ReplyToMessage string
}

type GoferEditedMessage struct {
	From string
	Message string
	OldMessage string
	Origin string
}
