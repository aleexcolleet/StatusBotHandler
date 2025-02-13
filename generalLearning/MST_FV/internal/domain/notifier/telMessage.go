package notifier

type Bot struct {
	APIToken string
	URLAPI   string
}
type Chats struct {
	ChatsId []string
}

func (bot *Bot) SetUrlApi() {
	bot.URLAPI = "https://api.telegram.org/" + bot.APIToken + "/sendMessage"
}
