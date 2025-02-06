package users

import (
	"cmd/main.go/internal/repositories"
	"context"
	"fmt"
)

/*
	InMemory.go is an adapter that implements userRepository interface
	In this case is an implementation on a in-Memory storage, no database, ...

ImUserStore [struct] -> represents the in-memory storage adapter.
In this case, it's an slice of users and
CreateUser and ReadUser implements the UserRepository interface. They need to
use repositories.User to implement it's package and refer to User struct.
*/
type ImUserStore struct {
	Users []repositories.User
}

// NewImUserStore is a constructor por a ImUserStore instance.
func NewImUserStore(ctx context.Context) *ImUserStore {
	return &ImUserStore{}
}

func (S *ImUserStore) CreateUser(ctx context.Context, user repositories.User) error {
	S.Users = append(S.Users, user)
	return nil
}

func (S *ImUserStore) ReadUser(ctx context.Context, userId int) (repositories.User, error) {
	for _, u := range S.Users {
		if userId == u.Id {
			return u, nil
		}
	}
	return repositories.User{}, fmt.Errorf("user not found")
}
