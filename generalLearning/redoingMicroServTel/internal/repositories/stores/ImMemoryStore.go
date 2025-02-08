package stores

import (
	"context"
	"redoingMicroServTel/internal/repositories"
)

type ImStore struct {
	ImURLsStore repositories.URLs
}

func NewImURLsStore(ctx context.Context) *ImStore {
	return &ImStore{
		ImURLsStore: repositories.URLs{},
	}
}

func (S *ImStore) loadUrl(ctx context.Context, urls repositories.URLs) error {
	S.ImURLsStore.URLs = urls.URLs
	return nil
}
