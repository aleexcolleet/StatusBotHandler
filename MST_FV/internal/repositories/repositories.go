package repositories

import (
	"context"
)

type URLs struct {
	URLs []string
}

type URLData struct {
	URL        string
	Status     bool
	Comments   string
	StatusCode int
}

type URLRepo interface {
	LoadURL(ctx context.Context, urls URLs) error
	GetURL(ctx context.Context) (URLs, error)
	LoadResponse(ctx context.Context, urlsResponse []URLData) error
}
