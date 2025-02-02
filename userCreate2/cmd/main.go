package main

import (
	"cmd/main.go/internal/domain/usecases"
	"cmd/main.go/internal/repositories/users"
	"context"
	"fmt"
)

/*
This application layer orchestrates the interaction between the domain and the adapters.
It's purpose is to initialize the adapter and the domain, and runs the application.
*/

func main() {
	UserStore := users.NewImUserStore(context.Background())
	Domain := usecases.NewDomain(context.Background(), UserStore)

	Domain.CreateUser()
	Domain.CreateUser()
	Domain.CreateUser()

	fmt.Printf("User1: %v\n", UserStore.Users[0])
	fmt.Printf("User2: %v\n", UserStore.Users[1])
	fmt.Printf("User3: %v\n", UserStore.Users[2])
}
