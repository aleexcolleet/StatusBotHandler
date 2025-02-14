package repositories

import (
	"MST_FV/internal/domain/models"
	"context"
)

//Port definitions

// UrlRepo interface
type UrlRepo interface {
	LoadUrls(ctx context.Context, urls models.URLs) error
	GetUrls(ctx context.Context) (models.URLs, error)
	LoadStatusResponse(ctx context.Context, urlsData []models.URLData) error
	GetStatusResponse(ctx context.Context) ([]models.URLData, error)
}

type CheckUrlRepo interface {
	GetCheckResp(ctx context.Context, urls models.URLs) ([]models.URLData, error)
}

// Message interface
type Message interface {
	GetMessages(ctx context.Context, urlsData []models.URLData) ([]string, error)
	// SendMessages
	//TODO I don't know why it's a messages []string
	SendMessages(ctx context.Context, messages []string) error
}
