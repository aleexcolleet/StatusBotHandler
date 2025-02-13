package repositories

import (
	"MST_FV2/internal/domain/models"
	"context"
)

type UrlRepo interface {
	LoadUrls(ctx context.Context, urls models.Urls) error
	GetUrls(ctx context.Context) (models.Urls, error)
	//SaveUrlsStatus(ctx context.Context, urlsData []models.UrlData) error
	//GetUrlsStatus(ctx context.Context) ([]models.UrlData, error)
}

type MessageRepo interface {
	GetMessage(ctx context.Context, urlsData []models.UrlData) ([]string, error)
	SendMessage(ctx context.Context, message []string) error
}
