package stores

import (
	"MST_FV2/configur"
	"MST_FV2/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type JsonStore struct {
	jsonStore     models.Urls
	jsonRouteRepo string
}

// NewJsonStore is a constructor for the Json implementation repo
func NewJsonStore(cfg configur.Config) *JsonStore {
	return &JsonStore{
		jsonStore:     models.Urls{},
		jsonRouteRepo: cfg.JsonFiles.JsonRepoRoute,
	}
}

func (s *JsonStore) LoadUrls(ctx context.Context, urls models.Urls) error {

	//urls.Urls... let me add the array directly without a loop
	var jsonData models.Urls
	jsonData.Urls = append(jsonData.Urls, urls.Urls...)

	updatedFile, err := json.MarshalIndent(jsonData, "", "-->") // Convierte la estructura en json y retorna el json
	if err != nil {
		return fmt.Errorf("Error al convertir a JSON: %w", err)
	}

	err = os.WriteFile(s.jsonRouteRepo, updatedFile, 0644) // Sobrescribe el archivo json con los datos json
	if err != nil {
		return fmt.Errorf("Error al escribir JSON: %w", err)
	}
	return nil
}

func (s *JsonStore) GetUrls(ctx context.Context) (models.Urls, error) {

	RepoFile, err := os.ReadFile(s.jsonRouteRepo)
	if err != nil {
		fmt.Println("Error reading JsonRepoFile in GetUrl")
		return models.Urls{}, err
	}

	var jsonData models.Urls
	err = json.Unmarshal(RepoFile, &jsonData)
	if err != nil {
		fmt.Println("Error UnMashalling JsonRepoFile in JsonData struct")
		return models.Urls{}, err
	}
	return jsonData, nil
}
