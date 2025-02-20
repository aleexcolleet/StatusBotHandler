package usecases

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"MST_FV/internal/repositories"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Services is the dependencies injection. Has the Port implementations to
// be used in the core part of the program
type Services struct {
	urlRepo      repositories.UrlRepo
	checkUrlRepo repositories.CheckUrlRepo
	messageRepo  repositories.Message
}

// ConsultAndSend Domain interface to interact with the adapters
type ConsultAndSend interface {
	consultAndSend(ctx context.Context, cfg config.Config) error
}

// NewServices is a constructor for the Services
func NewServices(urlRepo repositories.UrlRepo, messageRepo repositories.Message, checkUrlRepo repositories.CheckUrlRepo) *Services {
	return &Services{urlRepo: urlRepo, messageRepo: messageRepo, checkUrlRepo: checkUrlRepo}
}

func (s *Services) ConsultAndSend(ctx context.Context, cfg config.Config) error {

	err := s.loadUrls(ctx, cfg)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	urls, err := s.urlRepo.GetUrls(context.Background())
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	urlsDataResp, err := s.checkAndProcessUrls(context.Background(), urls)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	urls.UrlsData = urlsDataResp
	err = s.urlRepo.LoadStatusResponse(context.Background(), urls)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	msgs, err := s.messageRepo.GetMessages(context.Background(), urlsDataResp)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = s.messageRepo.SendMessages(context.Background(), msgs)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// LoadUrls() is a business logic func because it works directly with data.
// No matter where you store the Urls and Responses, the urls source is always the
// Json File, so I can directly implement it here

func (s *Services) loadUrls(ctx context.Context, cfg config.Config) error {

	jsonSourceFile := cfg.JsonRoutes.JsonRouteUrlSource
	var tmpURLs models.URLs

	file, err := os.ReadFile(jsonSourceFile)
	if err != nil {
		return err
	}
	var jsonData struct {
		Urls []string `json:"urls"`
	}
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		return err
	}

	tmpURLs.Urls = jsonData.Urls
	tmpURLs.UrlsData = []models.URLData{}

	err = s.urlRepo.LoadUrls(ctx, tmpURLs)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) checkAndProcessUrls(ctx context.Context, urls models.URLs) ([]models.URLData, error) {
	var urlsData []models.URLData

	for _, url := range urls.Urls {
		statusCode, duration, err := s.checkUrlRepo.GetCheckResp(context.Background(), url)

		urlData := models.URLData{
			Url:        url,
			Status:     statusCode < 400,
			Comment:    generateComment(statusCode, duration, err),
			StatusCode: statusCode,
		}
		urlsData = append(urlsData, urlData)
	}
	return urlsData, nil
}

func generateComment(statusCode int, duration time.Duration, err error) string {

	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	if statusCode >= 400 && statusCode < 500 {
		return fmt.Sprintf("\"📜 Hear ye, hear ye! error with the request made by the client: %d", statusCode)
	}
	if statusCode >= 500 && statusCode < 600 {
		return fmt.Sprintf("\"📜 Hear ye, hear ye! The web hath fallen with status: %d", statusCode)
	}
	return fmt.Sprintf("🛰️ Connection stable. The web is holding strong. Time to resp: %s ", duration)
}
