package stores

import (
	"MicroServ2/internal/repositories"
	"context"
)

type ImStore struct {
	ImStore repositories.URLs
}

func NewImStore() *ImStore {
	return &ImStore{
		ImStore: repositories.URLs{},
	}
}

func (S *ImStore) LoadURL(ctx context.Context, urls repositories.URLs) error {
	S.ImStore.URLs = urls.URLs
	return nil
}

func (S *ImStore) GetURL(ctx context.Context) (repositories.URLs, error) {
	return repositories.URLs{
		URLs: S.ImStore.URLs,
	}, nil
}
