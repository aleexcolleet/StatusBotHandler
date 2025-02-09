package notifier

type Bot struct {
	ApiToken string
	ApiUrl   string
}

type Chat struct {
	ChatId []string
}

// set API url TODO
// setURLAPI
func (B *Bot) SetUrlAPI() {
	B.ApiUrl = "https://api.telegram.org/bot" + B.ApiToken + "/sendMessage"
}
