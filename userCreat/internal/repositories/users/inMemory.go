package users

import (
	"cmd/main.go/internal/repositories"
	"context"
	"fmt"
)

/*
	InMemory database purpose is to work as an adapter. They're both the structure and methods that implements
	in a concrete way the ports from repositories.go. With this distinction I can easily replace the
	external systems without affecting the domain.

	Particulary, this code stores user data in-memory instead of a database.
*/

// ImUserStore is a user storage adapter. It basically
// stores users in a slice (simulating a database)
type ImUserStore struct {
	Users []repositories.User
}

// NewImUserStore is a constructor. Initializes a new instance of ImUserStore, which
// is an empty slice of users
func NewImUserStore(ctx context.Context) *ImUserStore {
	return &ImUserStore{
		Users: []repositories.User{},
	}
}

// CreateUser is a function to add a new User to the slice.
// In this case I can send repositories.UserRepo by value (not a poiner), because
// I don't need it. Just a copy to change the value of Users (which is actually changing)
func (S *ImUserStore) CreateUser(ctx context.Context, user repositories.User) error {
	S.Users = append(S.Users, user)
	return nil
}

// ReadUser method returns an instance of User by value. According to its Id
func (S *ImUserStore) ReadUser(ctx context.Context, userId int) (repositories.User, error) {
	for _, u := range S.Users {
		if u.Id == userId {
			return u, nil
		}
	}
	return repositories.User{}, fmt.Errorf("Cannot find user with id: %d", userId)
}
