package usecases

import (
	"context"
	"redoingMicroServTel/config"
	"redoingMicroServTel/internal/repositories"
)

type ImMemoryDomain struct {
	repo repositories.URLRepo
}

func NewImMemoryDomain(ctx context.Context, repo repositories.URLRepo) (*ImMemoryDomain, error) {
	return &ImMemoryDomain{
		repo: repo,
	}, nil
}

func (S *ImMemoryDomain) loadUrls(cfg config.Config) error {

	//S.repo.loadURLs(context.Background())
}
