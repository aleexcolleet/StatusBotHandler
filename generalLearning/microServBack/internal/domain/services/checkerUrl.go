package services

import (
	"context"
	"fmt"
	"microServBack/internal/repositories"
	"net/http"
	"time"
)

// UrlRepo is an instance of the repository to comunicate with him.
// This way I can use function like GetUrls and then store new data (responses),
// into the repo directly
type UrlRepo struct {
	repo repositories.UrlRepo
}

// NewCheckerAdapt si simply a constructor that receives the chosen repo
func NewCheckerAdapt(ctx context.Context, repo repositories.UrlRepo) *UrlRepo {
	return &UrlRepo{
		repo: repo,
	}
}

/*
	We need:
		a method-> checkUrls: that will receive a [repositories.URLs]
		struct and make the request with it, and then return a URLData
		struct with the new response data
		Then a LoadRequests func will acquire the data from repo,
		call CheckUrls with a loop and store values in a tmp [repositories.URLData]
		struct with a new func for the interface.
*/

func (receiver *UrlRepo) GetURLsStatus(ctx context.Context) error {
	//Fetch URLs from repository
	URLs, err := receiver.repo.GetUrls(context.Background())
	if err != nil {
		return fmt.Errorf("can't load URLs from repository: %w", err)
	}

	//Iterate across URLs from URLs struct to check em one by one and store resp in
	//tmpURLData
	var tmpURLData []repositories.URLData
	for _, urlReq := range URLs.URLs {
		urlResp, err := checkUrls(context.Background(), urlReq)
		if err != nil {
			return fmt.Errorf("error checking URLs with checkUrls: %w", err)
		}
		tmpURLData = append(tmpURLData, urlResp)
	}
	err = receiver.repo.StoreUrlsResp(context.Background(), tmpURLData)
	if err != nil {
		return fmt.Errorf("error storing URLs in repository: %w", err)
	}
	return nil
}

// First letter can't be a capital letter so that I can only call the func
// from this file
func checkUrls(ctx context.Context, URL string) (repositories.URLData, error) {

	//Use the timer to return if status is ok
	initTime := time.Now()

	//Make the request
	res, err := http.Get(URL)
	if err != nil {
		return repositories.URLData{}, fmt.Errorf("error making GET request: %w", err)
	}
	//Defer the body response allocation
	defer res.Body.Close()

	//Return a struct according to the res.code
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return repositories.URLData{
			URL:        URL,
			Status:     false,
			Comment:    "Caído ❌ HTTP error " + res.Status,
			StatusCode: res.StatusCode,
		}, nil
	}
	if res.StatusCode >= 500 && res.StatusCode < 600 {
		return repositories.URLData{
			URL:        URL,
			Status:     false,
			Comment:    "Caído ❌ HTTP error " + res.Status,
			StatusCode: res.StatusCode,
		}, nil
	}

	//if there isn't an error, check time and send resp statusa as active
	since := time.Since(initTime)
	return repositories.URLData{
		URL:     URL,
		Status:  true,
		Comment: "Active ✅ Response: " + since.String(),
	}, nil
}
