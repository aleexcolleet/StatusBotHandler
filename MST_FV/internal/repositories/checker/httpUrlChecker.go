package checker

import (
	"MST_FV/internal/domain/models"
	"context"
	"fmt"
	"net/http"
	"time"
)

// HttpUrlChecker is empty but needful. If I want to amplify it, it will be easy.
// I consider checkUrl a "static" func so that I only call it from the http implementation.
type HttpUrlChecker struct{}

func NewHttpUrlChecker() *HttpUrlChecker {
	return &HttpUrlChecker{}
}

// GetCheckResp receives the urls and returns the resp it gets from http request
func (s *HttpUrlChecker) GetCheckResp(ctx context.Context, urls models.URLs) ([]models.URLData, error) {

	var tmpUrlsData []models.URLData

	for _, tmpUrl := range urls.Urls {
		urlDataResp, err := checkUrl(tmpUrl)
		if err != nil {
			return []models.URLData{}, err
		}
		tmpUrlsData = append(tmpUrlsData, urlDataResp)
	}
	return tmpUrlsData, nil
}

func checkUrl(url string) (models.URLData, error) {

	initTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return models.URLData{}, fmt.Errorf("error making get request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return models.URLData{
			Url:        url,
			Status:     false,
			Comment:    "\"📜 Hear ye, hear ye! The web hath fallen with status: " + resp.Status,
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("error in statusCode: %s", url)
	}
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		return models.URLData{
			Url:        url,
			Status:     false,
			Comment:    "\"📜 Hear ye, hear ye! The web hath fallen with status: " + resp.Status,
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("error in statusCode: %s", url)

	}
	timeSince := time.Since(initTime)

	return models.URLData{
		Url:        url,
		Status:     true,
		Comment:    "🛰️ Connection stable. The web is holding strong. Time to resp: " + timeSince.String(),
		StatusCode: resp.StatusCode,
	}, nil
}
