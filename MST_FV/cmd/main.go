package main

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"MST_FV/internal/repositories/checker"
	"MST_FV/internal/repositories/message"
	"MST_FV/internal/repositories/stores"
	"context"
	"fmt"
)

/*
MicroService -> Telegram Message Requester

Things to do:

- Be more specific with errors for a better debug
- Use context.Context (to improve cancellations and deadline)
- The code could use goroutines to manage requests simultaneously
- Mock example for tests

*/

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Errorf("Error getting config: %v", err)
	}
	fmt.Printf("Config Creada\n")

	JsonStoreRepo := stores.NewJsonStoreRepo(cfg)
	fmt.Printf("Repo Creado con JsonStoreRepo\n")

	TelegramMsgs, err := message.NewTelegramMsgs(cfg)
	fmt.Printf("Repo Creado con TelegramMsgs\n")

	HttpChecker := checker.NewHttpUrlChecker()
	fmt.Printf("Repo Creado con HttpUrlChecker\n")

	Services := usecases.NewServices(JsonStoreRepo, TelegramMsgs, HttpChecker)
	fmt.Printf("Domain Creado con Services\n")

	Services.ConsultAndSend(context.Background())

}
