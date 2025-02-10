package stores

import (
	"context"
	"microServBack/internal/repositories"
)

type ImStore struct {
	UrlsStore repositories.URLs
}

func NewImStore() *ImStore {
	return &ImStore{
		UrlsStore: repositories.URLs{},
	}
}

// ImStore implementation of loading URLs for my database in inMemory
func (S *ImStore) LoadUrl(ctx context.Context, urls repositories.URLs) error {
	S.UrlsStore.URLs = urls.URLs
	return nil
}

// GetUrls in inmemory is an adaptation to fetch Urls from the repo
func (S *ImStore) GetUrls(ctx context.Context) (repositories.URLs, error) {
	return S.UrlsStore, nil
}
