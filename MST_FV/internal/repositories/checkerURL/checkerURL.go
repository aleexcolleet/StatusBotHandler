package checkerURL

import "C"
import (
	"MicroServ2/internal/repositories"
	"MicroServ2/internal/repositories/models"
	"context"
	"fmt"
	"net/http"
	"time"
)

// CheckerRepository CheckerRepo is an injection of dependencies. The actual repository
type CheckerRepository struct {
	urlRepo repositories.URLRepo
}

func newCheckerRepo(urlRepo repositories.URLRepo) *CheckerRepository {
	return &CheckerRepository{urlRepo: urlRepo}
}

func (C *CheckerRepository) CheckURLStatus(ctx context.Context) ([]models.URLData, error) {
	URLs, err := C.urlRepo.GetURL(context.Background())
	if err != nil {
		return []models.URLData{}, fmt.Errorf("error while fetching URLs: %v", err.Error())
	}

	var urlData []models.URLData
	for _, tmpUrl := range URLs.URLs {
		urlResp, err := checkerURL(context.Background(), tmpUrl)
		if err != nil {
			return []models.URLData{}, err
		}
		urlData = append(urlData, urlResp)
	}
	return urlData, nil
}

// checkerURL is a function that can only be called from this file.
func checkerURL(ctx context.Context, URL string) (models.URLData, error) {
	initTime := time.Now()
	res, err := http.Get(URL)
	if err != nil {
		return models.URLData{}, fmt.Errorf("error making the get fetching URL: %s", err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return models.URLData{
			URL:        URL,
			Status:     false,
			Comment:    "Caído ❌ HTTP error " + res.Status,
			StatusCode: res.StatusCode,
		}, nil
	}
	if res.StatusCode >= 500 && res.StatusCode <= 599 {
		return models.URLData{
			URL:        URL,
			Status:     false,
			Comment:    "Fallen ❌ HTTP error " + res.Status,
			StatusCode: res.StatusCode,
		}, nil
	}
	// In case the request worked
	since := time.Since(initTime)
	return models.URLData{
		URL:        URL,
		Status:     true,
		Comment:    "Active ✅ Response: " + since.String(),
		StatusCode: res.StatusCode,
	}, nil

}
