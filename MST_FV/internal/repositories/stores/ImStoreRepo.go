package stores

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"context"
)

type ImStoreRepo struct {
	urls models.URLs
	cfg  config.Config
}

func NewImStoreRepo(cfg config.Config) *ImStoreRepo {
	return &ImStoreRepo{
		cfg:  cfg,
		urls: models.URLs{},
	}
}

func (s *ImStoreRepo) LoadUrls(ctx context.Context, urls models.URLs) error {
	s.urls = urls
	return nil
}
func (s *ImStoreRepo) GetUrls(ctx context.Context) (models.URLs, error) {
	return s.urls, nil
}
func (s *ImStoreRepo) LoadStatusResponse(ctx context.Context, urls models.URLs) error {
	s.urls = urls
	return nil
}

func (s *ImStoreRepo) GetStatusResponse(ctx context.Context) ([]models.URLData, error) {
	return s.urls.UrlsData, nil
}
