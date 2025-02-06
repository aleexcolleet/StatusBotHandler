package services

import (
	"MicroServ2/internal/repositories"
	"context"
)

type CheckerURL struct {
	repo repositories.URLRepo
}

func NewCheckerURL(ctx context.Context, repo repositories.URLRepo) *CheckerURL {
	return &CheckerURL{
		repo: repo,
	}
}
