package stores

import (
	"context"
	"microServBack/internal/repositories"
)

type ImStore struct {
	UrlsStore    repositories.URLs
	UrlsDataResp []repositories.URLData
}

func NewImStore() *ImStore {
	return &ImStore{
		UrlsStore:    repositories.URLs{},
		UrlsDataResp: []repositories.URLData{},
	}
}

// LoadUrl implementation for my database in inMemory
func (S *ImStore) LoadUrl(ctx context.Context, urls repositories.URLs) error {
	S.UrlsStore.URLs = urls.URLs
	return nil
}

// GetUrls in inmemory is an adaptation to fetch Urls from the repo
func (S *ImStore) GetUrls(ctx context.Context) (repositories.URLs, error) {
	return S.UrlsStore, nil
}

// StoreUrlsResp is an adaptation to store URLs response from the Get request into the repository ImMemory
func (S *ImStore) StoreUrlsResp(ctx context.Context, UrlsData []repositories.URLData) error {
	S.UrlsDataResp = UrlsData
	return nil
}
