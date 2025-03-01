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
	jsonUrls       models.URLs
	jsonRouteRepo  string
	jsonSourceRepo string
}

func NewJsonStoreRepo(cfg config.Config) (*JsonStoreRepo, error) {
	return &JsonStoreRepo{
		jsonUrls:       models.URLs{},
		jsonSourceRepo: cfg.JsonRoutes.JsonRouteUrlSource,
		jsonRouteRepo:  cfg.JsonRoutes.JsonRouteRepo,
	}, nil
}

// LoadUrls loads the models.URLs to store em in the repo struct, in this case, a JsonFile
// In reality it's useless for my project but could come in handy, so I'll keep it
func (s *JsonStoreRepo) LoadUrls(ctx context.Context, urls models.URLs) error {

	//Transform the array into a []byte for Json recognition
	urlsInBytes, err := json.Marshal(urls)
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

// LoadStatusResponse should write the URL response into the Json file
func (s *JsonStoreRepo) LoadStatusResponse(ctx context.Context, urls models.URLs) error {

	urlsInBytes, err := json.MarshalIndent(urls, "", " ")
	if err != nil {
		return fmt.Errorf("error Marshaling urls: %s", err.Error())
	}
	err = os.WriteFile(s.jsonRouteRepo, urlsInBytes, 0644)
	if err != nil {
		return fmt.Errorf("error while writing into JsonFile: %s", err.Error())
	}
	return nil

}

func (s *JsonStoreRepo) GetStatusResponse(ctx context.Context) ([]models.URLData, error) {
	var tmpUrlsData []models.URLData
	file, err := os.ReadFile(s.jsonRouteRepo)
	if err != nil {
		return nil, fmt.Errorf("error opening JsonFile: %s", err.Error())
	}
	err = json.Unmarshal(file, &tmpUrlsData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JsonFile: %s", err.Error())
	}

	return tmpUrlsData, nil
}
