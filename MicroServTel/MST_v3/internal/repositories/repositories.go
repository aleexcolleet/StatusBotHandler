package repositories

import (
	"context"
)

/*
On this package we define the structs and interfaces that the program will use to communicate
adapters and domain.

We need a URL with its properties, and then a slice of URLs to store and access them.
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

// UserRepo is the interface that both repos will need to implement.
// LoadURLs will get URLs from wherever the repos take, and GetURLs will
// actually return an URLs struct with all values stored.
type UserRepo interface {
	LoadURLs(ctx context.Context) error
	GetURLs(ctx context.Context) (URLs, error)
}
