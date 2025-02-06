package usecases

import (
	"cmd/main.go/internal/repositories"
	"context"
	"fmt"
)

/*
	Domain layer contains the core business logic and defines the ports (interfaces)
	that external systems must implement.

	In this case we define the logic for user management.
	Users struct: Is a receiver of the interface. Represents de domain logic and
	depends on the UserRepo.
*/

type Users struct {
	Repo repositories.UserRepository
}

// NewDomain constructor injects de repository (adapter) into the domain.
func NewDomain(ctx context.Context, repo repositories.UserRepository) *Users {
	return &Users{
		Repo: repo,
	}
}

/*
CreateUser is a logic implementation of the bussines.
It doesn't care how the information is stored, it just needs to send it the way
the interface says as a contract.
1. User is created with the specifications of the port.
2. CreateUser() is called with the user instance. We use the receiver Repo.
*/
func (receiver Users) CreateUser() error {
	user := repositories.User{"alex", 444, "holabuenas", "alex@gmail.com"}
	err := receiver.Repo.CreateUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
