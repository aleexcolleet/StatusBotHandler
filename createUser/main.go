package main

import (
	"context"
	"createUser/internal/domain/usecases"
	"createUser/internal/repositories/users"
	"fmt"
)

func main() {
	UserStore := users.NewlmUserStore(context.Background())
	Domain := usecases.NewDomain(context.Background(), UserStore)
	Domain.CreateUser()
}
