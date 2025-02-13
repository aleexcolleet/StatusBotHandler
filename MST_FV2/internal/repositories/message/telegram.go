package message

import (
	"MST_FV2/configur"
	"MST_FV2/internal/domain/models"
	"context"
)

type Telegram struct {
	bot     Bot
	chatsId ChatsId
}

type Bot struct {
	apiToken string
	urlApi   string
}

type ChatsId struct {
	chatId []string
}

func NewTelegram(cfg configur.Config) *Telegram {

	botLoad := Bot{
		apiToken: cfg.Bot.APIToken,
	}
	botLoad.setUrlApi()

	chatLoad := ChatsId{
		chatId: cfg.Chats.ChatsId,
	}

	return &Telegram{
		bot:     botLoad,
		chatsId: chatLoad,
	}
}
func (b *Bot) setUrlApi() {
	b.urlApi = "https://api.telegram.org/bot" + b.apiToken + "/sendMessage"
}

// TODO implementation of these two
func (t *Telegram) GetMessage(ctx context.Context, urlsData []models.UrlData) ([]string, error) {
	return []string{}, nil
}
func (t *Telegram) SendMessage(ctx context.Context, message []string) error {
	return nil
}
