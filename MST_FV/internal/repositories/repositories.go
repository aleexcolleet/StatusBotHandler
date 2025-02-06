package repositories

import (
	"context"
)

type URLs struct {
	URLs []string
}

type URLRepo interface {
	LoadURL(ctx context.Context, urls URLs) error
	GetURL(ctx context.Context) (URLs, error)
	//
	//
}
