package repositories

import (
	"context"
)

type URL struct {
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
}
