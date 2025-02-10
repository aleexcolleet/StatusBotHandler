package services

import (
	"context"
	"microServBack/internal/repositories"
)

// UrlRepo is an instance of the repository to comunicate with him.
// This way I can use function like GetUrls and then store new data (responses),
// into the repo directly
type UrlRepo struct {
	repo repositories.UrlRepo
}

// NewUrlRepo si simply a constructor that receives the chosen repo
func NewUrlRepo(ctx context.Context, repo repositories.UrlRepo) *UrlRepo {
	return &UrlRepo{
		repo: repo,
	}
}

/*
TODO
	We need:
		a method-> checkUrls: that will receive a [repositories.URLs]
		struct and make the request with it, and then return a URLData
		struct with the new response data
		Then a LoadRequests func will acquire the data from repo,
		call CheckUrls with a loop and store values in a tmp [repositories.URLData]
		struct. That way I can implement a new func in the repo to store values.
*/

//Todo
func (receiver *UrlRepo) LoadRequests(ctx context.Context) error {

}

// First letter can't be a capital letter so that I can only call the func
// from this file
func checkUrls(ctx context.Context, URLs string) (repositories.URLData, error) {
	var tmpURLData repositories.URLData
	_ = tmpURLData
}
