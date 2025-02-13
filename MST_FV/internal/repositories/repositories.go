package repositories

import (
	"MST_FV/internal/domain/models"
	"context"
)

//Port definitions

//UrlRepo
type UrlRepo interface {
	LoadUrls(ctx context.Context, urls models.URLs) error
	GetUrls(ctx context.Context) (models.URLs, error)
	//LoadStatusResponse
	//GetStatusResponse
}

//Message
type Message interface {
	GetMessages(ctx context.Context, models.URLData) ([]string, error)

	// SendMessages
	//TODO I don't know why its a messages []string
	SendMessages(ctx context.Context, messages []string) error
}
