package message

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

type TelegramMsgs struct {
	messages       []TelegramMsg
	telegramBotApi TelegramBotApi
	chatIds        []string
	apiUrlMess     string
	urls           []string
}

// TelegramMsg is a struct to manage the specific message characteristics from telegram
type TelegramMsg struct {
	message  string
	msgIndex int
}
type TelegramBotApi struct {
	bot *tlgrmBotApi.BotAPI
}

func NewTelegramMsgs(cfg config.Config) (*TelegramMsgs, error) {

	// Create bot instance with the actual api Token
	bot, err := tlgrmBotApi.NewBotAPI(cfg.Bot.ApiToken)
	if err != nil {
		fmt.Errorf("error creating bot instance: %s", err)
	}

	chatsLoad := cfg.Chats.ChatsId   //Create the chatIds array
	apiUrlMess := cfg.Bot.ApiUrlMess //tokenUrl

	return &TelegramMsgs{

		chatIds:  chatsLoad,
		messages: []TelegramMsg{},
		telegramBotApi: TelegramBotApi{
			bot: bot,
		},
		apiUrlMess: apiUrlMess,
	}, nil
}

// GetMessages func converts the urlsData into a TelegramMsgs struct so that I can
// use it more easily here.
// I need to return a []slice and not a TelegramMessage type because this is just an implementation.
func (t *TelegramMsgs) GetMessages(ctx context.Context, urlsData []models.URLData) ([]string, error) {

	var tmpTelMsg TelegramMsg

	var tmpTelMsgs []TelegramMsg
	var tmpMessages []string
	// TelegramMsgs constr
	for i, tmpUrlsData := range urlsData {

		tmpTelMsg.msgIndex = i + 1                                                                            //iterator
		tmpTelMsg.message = fmt.Sprintf("[The url %d: on -----] %s", tmpTelMsg.msgIndex, tmpUrlsData.Comment) //telMess

		tmpTelMsgs = append(tmpTelMsgs, tmpTelMsg)           //messages struct
		tmpMessages = append(tmpMessages, tmpTelMsg.message) //[]string messages
	}
	t.messages = append(t.messages, tmpTelMsgs...)
	return tmpMessages, nil
}

// SendMessages converts the messages slice into the required struct by telegram(chatId + message)
// and sends it with a loop so that all chats have all Messages
// I find unnecessary to store the msg struct for telegram, so I'll just generate it and send it here.
func (t *TelegramMsgs) SendMessages(ctx context.Context, messages []string) error {

	apiUrlMess := t.apiUrlMess
	chatIds := t.chatIds

	for _, message := range messages {
		for _, chatId := range chatIds {

			reqBody, _ := json.Marshal(map[string]string{
				"chat_id": chatId,
				"text":    message,
			})
			resp, err := http.Post(apiUrlMess, "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				return fmt.Errorf("error sending message: %s", err)
			}
			defer resp.Body.Close() //keep it for debugging if needed

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error sending message: %s", resp.Status)
			}
		}
	}
	fmt.Printf("✅Messages sent correctly✅\n")
	return nil
}
