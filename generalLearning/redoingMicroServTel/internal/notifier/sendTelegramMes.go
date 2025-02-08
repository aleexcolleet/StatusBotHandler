package notifier

type Bot struct {
	ApiURL string
}
type Chats struct {
	ChatsId []string
}

func (B *Bot) SetApiURL() {
	B.ApiURL = "https://api.telegram.org/bot" + B.ApiURL + "/sendMessage"
}
