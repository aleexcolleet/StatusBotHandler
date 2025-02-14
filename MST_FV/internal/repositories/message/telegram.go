package message

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"context"
	"fmt"
	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramMsg is a struct to manage the specific message characteristics from telegram
type TelegramMsg struct {
	telBot   TelegramBotApi
	message  string
	msgIndex int
}
type TelegramMsgs []TelegramMsg

type TelegramBotApi struct {
	bot *tlgrmBotApi.BotAPI
}

func NewTelegramMsgs(cfg config.Config, urls models.URLs) *TelegramMsgs {
	bot, err := tlgrmBotApi.NewBotAPI(cfg.Bot.ApiToken)
	if err != nil {
		fmt.Errorf("error creating bot instance: %s", err)
	}
	return &TelegramMsgs{}
}

func (t *TelegramMsgs) GetMessages(ctx context.Context, urlsData []models.URLData) ([]string, error) {
	var tmpTelMsg TelegramMsg
	var tmpTelMsgs []TelegramMsg

	for i, urlData := range urlsData {
		tmpTelMsg.msgIndex = i + 1
		tmpTelMsg.message = urlData.Comment
		tmpTelMsgs = append(tmpTelMsgs, tmpTelMsg)
	}

}
