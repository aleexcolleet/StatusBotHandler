package users

import (
	"context"
	"createUser/internal/repositories"
	"fmt"
)

type lmUserStore struct {
	Users []repositories.User
}

func NewlmUserStore(ctx context.Context) *lmUserStore {
	return &lmUserStore{
		Users: []repositories.User{},
	}

func (S* lmUserStore) CreateUser(ctx context.Context, user repositories.User) error {
	S.Users = append(S.Users, user)
	return nil
	}
}
