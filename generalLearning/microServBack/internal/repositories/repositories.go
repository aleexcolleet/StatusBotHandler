package repositories

import (
	"context"
)

type URLData struct {
	URL        string
	Status     bool
	Comment    string
	StatusCode int
}

// URLs Struct that I'll use as a storer in inMemory
type URLs struct {
	URLs []string
}

type UrlRepo interface {
	LoadUrl(ctx context.Context, urls URLs) error
	GetUrls(ctx context.Context) (URLs, error)
	StoreUrlsResp(ctx context.Context, UrlsData []URLData) error
}
