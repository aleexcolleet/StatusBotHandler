package repositories

import "context"

/*
On this package we define the structs and interfaces that the program will use to communicate
adapters and domain.

*/

// URL is a struct that contains all information
type URL struct {
	Url string
	Id  int
}

// URLs is a struct with a slice of URLs
type URLs struct {
	URLs []URL
}

type UserRepo interface {
	LoadURLs(ctx context.Context) error
	GetURLs(ctx context.Context) (*URLs, error)
}
