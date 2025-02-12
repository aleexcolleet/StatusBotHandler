package stores

import (
	"MicroServ2/internal/repositories/models"
	"context"
)

type ImStore struct {
	ImStore     models.URLs
	ImStoreResp []models.URLData
}

func NewImStore() *ImStore {
	return &ImStore{
		ImStore:     models.URLs{},
		ImStoreResp: []models.URLData{},
	}
}

func (S *ImStore) LoadURL(ctx context.Context, urls models.URLs) error {
	S.ImStore.URLs = urls.URLs
	return nil
}

func (S *ImStore) GetURL(ctx context.Context) (models.URLs, error) {
	return models.URLs{
		URLs: S.ImStore.URLs,
	}, nil
}

func (S *ImStore) LoadResponse(ctx context.Context, urlsResponse []models.URLData) error {
	S.ImStoreResp = urlsResponse
	return nil
}
