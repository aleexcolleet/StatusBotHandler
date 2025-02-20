package repositories

import (
	"MST_FV/internal/domain/models"
	"context"
	"time"
)

//Port definitions

// UrlRepo interface
type UrlRepo interface {
	LoadUrls(ctx context.Context, urls models.URLs) error
	GetUrls(ctx context.Context) (models.URLs, error)
	LoadStatusResponse(ctx context.Context, urls models.URLs) error
	GetStatusResponse(ctx context.Context) ([]models.URLData, error)
}

// CheckUrlRepo interface
type CheckUrlRepo interface {
	GetCheckResp(ctx context.Context, url string) (int, time.Duration, error)
}

// Message interface
type Message interface {
	GetMessages(ctx context.Context, urlsData []models.URLData) ([]string, error)
	SendMessages(ctx context.Context, messages []string) error
}
