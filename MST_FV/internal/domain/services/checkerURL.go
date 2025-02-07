package services

import (
	"MicroServ2/internal/repositories"
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
	checkerURL.go is a service that I use to request httpUrls from a client
	and returns it to store the response and other data into the repo.
	HOW?

	- Create the constructor for an instance of the port (repo interface)
	- Fetch URLs struct from storage in repo.
	- Check URls with CheckURL one by one so that I can treat and store errors.
		This func returns a tmp struct repositories.URLData
	-
*/

type CheckerURL struct {
	repo repositories.URLRepo
}

func NewCheckerURL(ctx context.Context, repo repositories.URLRepo) *CheckerURL {
	return &CheckerURL{
		repo: repo,
	}
}

func (C *CheckerURL) GetURLStatus(ctx context.Context) error {
	// Fetch URLs from repo
	URLs, err := C.repo.GetURL(ctx)
	if err != nil {
		return fmt.Errorf("error fetching URLs from repo: %s", err.Error())
	}

	//Iterate across URLs to check its response and append resp struct
	//to the tmp struct
	var urlData []repositories.URLData
	for _, tmpUrl := range URLs.URLs {
		urlResp, err := CheckURL(context.Background(), tmpUrl)
		if err != nil {
			return fmt.Errorf("error checking URL: %s", err.Error())
		}
		urlData = append(urlData, urlResp)
	}
	C.repo.LoadResponse(context.Background(), urlData)
	return nil
}

func CheckURL(ctx context.Context, URL string) (repositories.URLData, error) {
	// We use a timer to know how long it takes the URL to be requested
	initTime := time.Now()
	// Request the URL with a GET method
	res, err := http.Get(URL)
	// err != woudl mean something went wrong before even receiving the response.
	//fallen URL, incorrect [*DNS], timeout, etc...
	if err != nil {
		return repositories.URLData{}, fmt.Errorf("error requesting URL: %s", err.Error())
	}
	// This grants that the Body resp is closed at the end of the func no matter what.
	defer res.Body.Close()

	//Status Handling
	if res.StatusCode >= 400 && res.StatusCode <= 499 {
		return repositories.URLData{
			URL:        URL,
			Status:     false,
			Comments:   "Caído ❌ HTTP error " + res.Status,
			StatusCode: res.StatusCode,
		}, nil
	}
	if res.StatusCode >= 500 && res.StatusCode <= 599 {
		return repositories.URLData{
			URL:        URL,
			Status:     false,
			Comments:   "Fallen ❌ HTTP error " + res.Status,
			StatusCode: res.StatusCode,
		}, nil
	}
	// If there isn't and error, we send Active message + elapsed time
	since := time.Since(initTime)
	return repositories.URLData{
		URL:        URL,
		Status:     true,
		Comments:   "Active ✅ Response: " + since.String(),
		StatusCode: res.StatusCode,
	}, nil
}

/*
	*DNS: Translation between domain name and its IP.
		Could mean an error if translation fails
*/
