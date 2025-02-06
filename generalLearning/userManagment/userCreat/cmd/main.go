package main

import (
	"cmd/main.go/internal/domain/usecases"
	"cmd/main.go/internal/repositories/users"
	"context"
	"fmt"
)

/*
1. Create an UserStore(in-memory storage)
2. Passes it into the domain
3. Calls Domain.CreateUser()
4. Prints the stored user's name

Only orchestrates the logic, but doesn't store anything itself
*/
func main() {
	UserStore := users.NewImUserStore(context.Background())
	Domain := usecases.NewDomain(context.Background(), UserStore)
	Domain.CreateUser()
	fmt.Printf("User created: %v", UserStore.Users[0])
}
