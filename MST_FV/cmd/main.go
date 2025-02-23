package main

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"MST_FV/internal/repositories/checker"
	"MST_FV/internal/repositories/message"
	"MST_FV/internal/repositories/stores"
	"MST_FV/server"
	"log"
)

/*
MicroService -> Telegram Message Requester

TODO Things to improve project:
- Implement a server cape so that I can expose my business logic witt an HTTP API
	with the framework "fiber"

- Be more specific with errors for a better debug
- Use context.Context (to improve cancellations and deadline)
- The code could use goroutines to manage requests simultaneously
- Mock example for tests
- Use a logging system to register important events

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

	//init server
	server := server.NewServer(cfg, services)
	server.SetUpRoutes()

	log.Println("Starting server on 3000")
	err = server.Start()
	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
	log.Fatal("Server Started Successfully\n")
}
