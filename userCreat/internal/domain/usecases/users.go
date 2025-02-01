package usecases

import (
	"cmd/main.go/internal/repositories"
	"context"
	"fmt"
)

/*
Users struct is the Domain Layer that interacts with the repository.
Repo repositories.UserRepo is the injected repository (port)

The Users struct doesn't care whether the Repository is using a database,
an API, or in-memory storage.
That way, Domain Layer is decoupled from infrastructure (Repositories).
*/
type Users struct {
	Repo repositories.UserRepo
}

/*
NewDomain is a constructor that creates and initializes a Users struct
with an injected interface, which means we can pass any repository
implementation.
*/
func NewDomain(ctx context.Context, repo repositories.UserRepo) *Users {
	return &Users{
		Repo: repo,
	}
}

//CreateUser method

func (receiver Users) CreateUser() error {
	user := repositories.User{"alex", 4444, 12, "alex@gmail.com"}
	err := receiver.Repo.CreateUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Error while creating user: %v", err)
	}
	return nil
}
