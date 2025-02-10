package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"microServBack/config"
	"microServBack/internal/repositories"
	"os"
)

// ImMemoryStore Struct to implement the ImMemory repo interface
type ImMemoryStore struct {
	repo repositories.UrlRepo
}

type JsonFile struct {
	Urls []string `json:"urls"`
}

func NewImMemoryStore(ctx context.Context, repo repositories.UrlRepo) *ImMemoryStore {
	return &ImMemoryStore{
		repo: repo,
	}
}

// LoadUrl Implementation so that I can call the repo one with domain
func (receiver *ImMemoryStore) LoadUrl(cfg config.Config) error {

	//With my config struct, I can acquire the PATH to Json file.
	var jsonFile JsonFile
	file, err := os.ReadFile(cfg.JsonUrlRoute)
	if err != nil {
		return fmt.Errorf("error while loading JsonFileUrls: %w", err)
	}
	// Transform []byte into a []string and store it in the tmpStruct
	err = json.Unmarshal(file, &jsonFile)
	if err != nil {
		return fmt.Errorf("error while converting JsonFileUrls: %w", err)
	}

	urlRode := repositories.URLs{
		URLs: jsonFile.Urls,
	}
	//The idea is acquiring Urls to send them to the repo interface
	err = receiver.repo.LoadUrl(context.Background(), urlRode)
	return nil
}

func (receiver *ImMemoryStore) GetUrls(ctx context.Context) (repositories.URLs, error) {
	Urls, err := receiver.repo.GetUrls(ctx)
	if err != nil {
		return repositories.URLs{}, fmt.Errorf("error while getting Urls in domain: %w", err)
	}
	return Urls, nil
}
