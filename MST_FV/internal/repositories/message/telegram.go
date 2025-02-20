package message

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramMsgs struct {
	Bot   Bot
	Chats Chats
}
type Bot struct {
	ApiToken   string
	ApiUrlMess string
}

type Chats struct {
	ChatsId []string
}

func NewTelegramMsgs(cfg config.Config) *TelegramMsgs {

	botLoad := Bot{
		ApiToken: cfg.Bot.ApiToken,
	}
	botLoad.setUrlApi()

	chatsLoad := Chats{
		ChatsId: cfg.Chats.ChatsId,
	}
	return &TelegramMsgs{
		Bot:   botLoad,
		Chats: chatsLoad,
	}
}

// GetMessages func converts the urlsData into a TelegramMsgs struct so that I can
// use it more easily here.
// I need to return a []slice and not a TelegramMessage type because this is just an implementation.
func (t *TelegramMsgs) GetMessages(ctx context.Context, urlsData []models.URLData) ([]string, error) {

	var tmpString []string

	for i, urlData := range urlsData {
		tmpMesStr := fmt.Sprintf("Url number[%d] checked: [%s] \n%s\n", i, urlData.Url, urlData.Comment)
		tmpString = append(tmpString, tmpMesStr)
	}
	return tmpString, nil
}

// SendMessages converts the messages slice into the required struct by telegram(chatId + message)
// and sends it with a loop so that all chats have all Messages
// I find unnecessary to store the msg struct for telegram, so I'll just generate it and send it here.
func (t *TelegramMsgs) SendMessages(ctx context.Context, messages []string) error {

	botLoad := t.Bot
	chatLoad := t.Chats

	for _, msg := range messages {
		for _, chatId := range chatLoad.ChatsId {
			reqMap, err := json.Marshal(map[string]string{
				"chat_id": chatId,
				"text":    msg,
			})
			resp, err := http.Post(botLoad.ApiUrlMess, "application/json", bytes.NewBuffer(reqMap))
			if err != nil {
				return fmt.Errorf("error sending message: %s", err)
			}
			resp.Body.Close() //keep it for debugging if needed
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error sending message: %s", resp.Status)
			}
		}
	}
	fmt.Printf("✅Messages sent correctly✅\n")
	return nil
}

// Dependencies
func (b *Bot) setUrlApi() {
	b.ApiUrlMess = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.ApiToken)
}
