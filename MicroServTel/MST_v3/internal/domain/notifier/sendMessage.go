package notifier

type Bot struct {
	ApiToken string
}

type Chat struct {
	ChatId string
}
type Chats struct {
	ChatId []Chat
}
