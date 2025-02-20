package checker

import (
	"context"
	"net/http"
	"time"
)

// HttpUrlChecker is empty but needful. If I want to amplify it, it will be easy.
// I consider checkUrl a "static" func so that I only call it from the http implementation.
type HttpUrlChecker struct{}

func NewHttpUrlChecker() *HttpUrlChecker {
	return &HttpUrlChecker{}
}

// GetCheckResp receives the urls and returns the raw resp, but It doesn't process the resp because
// that is business logic
func (s *HttpUrlChecker) GetCheckResp(ctx context.Context, url string) (int, time.Duration, error) {
	initTime := time.Now()
	resp, err := http.Get(url)
	timeSince := time.Since(initTime)
	if err != nil {
		return 0, timeSince, err
	}
	return resp.StatusCode, timeSince, nil
}
