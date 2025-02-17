package message

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"context"
	"fmt"
	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type TelegramMsgs struct {
	messages       []TelegramMsg
	telegramBotApi TelegramBotApi
	chatIds        []int64
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

	//Create the chatIds array converted to int64
	chatIds := make([]int64, len(cfg.Chats.ChatsId))
	for _, cId := range cfg.Chats.ChatsId {
		tmpChatId, err := strconv.ParseInt(cId, 10, 64)
		if err != nil {
			return &TelegramMsgs{}, fmt.Errorf("error parsing chatId: %s", err)
		}
		chatIds = append(chatIds, tmpChatId)
	}

	return &TelegramMsgs{
		chatIds:  chatIds,
		messages: []TelegramMsg{},
		telegramBotApi: TelegramBotApi{
			bot: bot,
		},
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

		tmpTelMsg.msgIndex = i + 1              //iterator
		tmpTelMsg.message = tmpUrlsData.Comment //messageToSend

		tmpTelMsgs = append(tmpTelMsgs, tmpTelMsg)             //messages struct
		tmpMessages = append(tmpMessages, tmpUrlsData.Comment) //[]string messages
	}
	t.messages = append(t.messages, tmpTelMsgs...)
	return tmpMessages, nil
}

// SendMessages converts the messages slice into the required struct by telegram(chatId + message)
// and sends it with a loop so that all chats have all Messages
// I find unnecessary to store the msg struct for telegram, so I'll just generate it and send it here.
func (t *TelegramMsgs) SendMessages(ctx context.Context, messages []string) error {

	for _, message := range messages {
		for _, chatId := range t.chatIds {
			msg := tlgrmBotApi.NewMessage(chatId, message)
			_, err := t.telegramBotApi.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("error while sending message to a Telegram chat: %v\n", err)
			}
		}
	}
	return nil
}
