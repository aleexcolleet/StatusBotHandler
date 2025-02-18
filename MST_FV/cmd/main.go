package main

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"MST_FV/internal/repositories/checker"
	"MST_FV/internal/repositories/message"
	"MST_FV/internal/repositories/stores"
	"context"
	"fmt"
	"log"
)

/*
MicroService -> Telegram Message Requester

things to do:
- every definition port goes into repositories.go

Things to do:

- Be more specific with errors for a better debug
- Use context.Context (to improve cancellations and deadline) TODO
- The code could use goroutines to manage requests simultaneously
- Mock example for tests

*/

type Dependencies struct {
	Config       config.Config
	JsonStore    *stores.JsonStoreRepo
	TelegramMsgs *message.TelegramMsgs
	HttpChecker  *checker.HttpUrlChecker
	Services     *usecases.Services
}

func newDependencyInjector() (*Dependencies, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting configLoad %v\n", err)
	}

	//ReposLoad
	jsonStoreRepo := stores.NewJsonStoreRepo(cfg)
	telegramMsgs, err := message.NewTelegramMsgs(cfg)

	if err != nil {
		return nil, fmt.Errorf("error creating telegramMsgs %v\n", err)
	}
	httpChecker := checker.NewHttpUrlChecker()
	services := usecases.NewServices(jsonStoreRepo, telegramMsgs, httpChecker)

	return &Dependencies{
		Config:       cfg,
		JsonStore:    jsonStoreRepo,
		TelegramMsgs: telegramMsgs,
		HttpChecker:  httpChecker,
		Services:     services,
	}, nil
}

func main() {

	injector, err := newDependencyInjector()
	if err != nil {
		log.Fatalf("error creating injector: %v", err)
	}

	err = injector.Services.ConsultAndSend(context.Background())
	if err != nil {
		log.Fatalf("error in consultAndSend services: %v", err)
	}
	log.Println("ConsultAndSend success")
}
