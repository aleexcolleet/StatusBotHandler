package main

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"MST_FV/internal/repositories/checker"
	"MST_FV/internal/repositories/message"
	"MST_FV/internal/repositories/stores"
	"context"
	"log"
	"time"
)

/*
MicroService -> Telegram Message Requester

Things to do:

- Be more specific with errors for a better debug
- Use context.Context (to improve cancellations and deadline) TODO
- The code could use goroutines to manage requests simultaneously
- Mock example for tests

*/

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error getting configLoad %v\n", err)
	}

	// Adapters
	storeRepo := stores.NewJsonStoreRepo(cfg)
	msgsRepo := message.NewTelegramMsgs(cfg)
	checkerRepo := checker.NewHttpUrlChecker()

	// Domain Service with dependencies injected
	services := usecases.NewServices(storeRepo, msgsRepo, checkerRepo)

	err = services.ConsultAndSend(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
}
