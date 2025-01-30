package usecases

import (
	"context"
	"createUser/internal/domain/usecases"
	"fmt"
)

type Users struct {
	Repo repositories.UserRepo
}

func NewDomain(ctx context.Context, repo repositories.UserRepo) *Users {
	&Users{
		Repo: repo,
	}
}
