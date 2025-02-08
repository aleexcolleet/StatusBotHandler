package repositories

import "context"

type URLs struct {
	URLs []URL
}

type URL struct {
	URL        string
	Status     bool
	Comment    string
	StatusCode int
}

type URLRepo interface {
	loadUrls(ctx context.Context, urls URLs) error
}
