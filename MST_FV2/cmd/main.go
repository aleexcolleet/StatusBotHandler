package main

import (
	"MST_FV2/configur"
	"MST_FV2/internal/domain"
	"MST_FV2/internal/repositories/message"
	"MST_FV2/internal/repositories/stores"
	"fmt"
)

/*
	Devop Checker

	telegram repo adaptation

	imMemory adaptation

	create a domain to call load urls, ...

*/

func main() {
	cfg, err := configur.GetConfig()
	if err != nil {
		fmt.Errorf("error fetching configur: %v", err)
	}

	JsonRepo := stores.NewJsonStore(cfg)
	//whatsapp := newwhatsapp(cfg)
	Telegram := message.NewTelegram(cfg)

	Services := domain.NewServices(JsonRepo, Telegram)
	Services.AnalyzeAndSend()
}
