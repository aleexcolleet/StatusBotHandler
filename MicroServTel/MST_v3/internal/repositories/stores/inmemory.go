package stores

import (
	"cmd/main.go/internal/repositories"
	"context"
)

// URLsToReceive contains the predefined URLs to return with GetURLs
var URLsToReceive = []string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"http://neverssl.com/",
}

/*
In-memory is the adapter to repositories which works with local database. In this case we
need to store the URLs that are already specified in this file.

First we need a struct to store URLs
Then a constructor for this struct
*/

// ImUserStore will store all URLs to return em as a []repositories.Url
type ImUserStore struct {
	URLs []repositories.URL
}

// NewImUserStore is a constructor that stores URLs consts
// in memory.
// Each round of the loop adds a new repositories.URL to the slice with the values from URLsToReceive
func NewImUserStore(ctx context.Context) (*ImUserStore, error) {
	var urls []repositories.URL
	for i, url := range URLsToReceive {
		urls = append(urls, repositories.URL{
			Url: url,
			Id:  i + 1,
		})
	}
	return &ImUserStore{URLs: urls}, nil
}

// LoadURLs port is adapted here according to URLs from the const values.
// I think that this func is not necessary on this adaptation
func (S *ImUserStore) LoadURLs(ctx context.Context) error { return nil }

// GetURLs simply returns the struct created on NewImUserStore
func (S *ImUserStore) GetURLs(ctx context.Context) (repositories.URLs, error) {
	return repositories.URLs{
		URLs: S.URLs,
	}, nil
}
