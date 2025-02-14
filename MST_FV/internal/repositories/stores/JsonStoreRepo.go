package stores

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type JsonStoreRepo struct {
	jsonUrls      models.URLs
	jsonRouteRepo string
}

func NewJsonStoreRepo(cfg config.Config) *JsonStoreRepo {
	return &JsonStoreRepo{
		jsonUrls:      models.URLs{},
		jsonRouteRepo: cfg.JsonRoutes.JsonRouteUrlSource,
	}
}

// LoadUrls loads the models.URLs to store em in the repo struct, in this case, a JsonFile
func (s *JsonStoreRepo) LoadUrls(ctx context.Context, urls models.URLs) error {

	//Transform the array into a []byte for Json recognition
	urlsInBytes, err := json.MarshalIndent(urls.Urls, "", "-->")
	if err != nil {
		return fmt.Errorf("error Marshaling urls: %s", err.Error())
	}
	//Then write into the Json file
	err = os.WriteFile(s.jsonRouteRepo, urlsInBytes, 0644)
	if err != nil {
		return fmt.Errorf("error while writting into JsonFile: %s", err.Error())
	}
	return nil
}
func (s *JsonStoreRepo) GetUrls(ctx context.Context) (models.URLs, error) {

	//tmp instance to return the urls
	var tmpUrls models.URLs

	//Os.ReadFile reads and closes the file automatically withoud defer need
	file, err := os.ReadFile(s.jsonRouteRepo)
	if err != nil {
		return models.URLs{}, fmt.Errorf("error opening JsonFile: %s", err.Error())
	}

	err = json.Unmarshal(file, &tmpUrls)
	if err != nil {
		return models.URLs{}, fmt.Errorf("error unmarshalling JsonFile: %s", err.Error())
	}

	return tmpUrls, nil
}

func (s *JsonStoreRepo) LoadStatusResponse(ctx context.Context, urlsData []models.URLData) error {
	s.jsonUrls.UrlsData = urlsData
	return nil
}

func (s *JsonStoreRepo) GetStatusResponse(ctx context.Context) ([]models.URLData, error) {
	return s.jsonUrls.UrlsData, nil
}
