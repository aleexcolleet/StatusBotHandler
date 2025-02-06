package usecases

import (
	"MicroServ2/config"
	"MicroServ2/internal/repositories"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type ImMemoryURL struct {
	repo repositories.URLRepo
}

func NewDomainInMemory(ctx context.Context, repo repositories.URLRepo) *ImMemoryURL {
	return &ImMemoryURL{repo: repo}
}

type JsonTmp struct {
	URLs []string `json:"urls"`
}

func (receiver *ImMemoryURL) LoadURL(cfg config.Config) error {
	file, err := os.ReadFile(cfg.JsonFile.JsonUrlRoute)
	if err != nil {
		return fmt.Errorf("Error while loading URL: %v", err)
	}

	var jsonTmp JsonTmp
	err = json.Unmarshal(file, &jsonTmp)
	if err != nil {
		return fmt.Errorf("Error while loading URL: %v", err)
	}

	urlRode := repositories.URLs{
		URLs: jsonTmp.URLs,
	}
	err = receiver.repo.LoadURL(context.Background(), urlRode)
	return nil
}

func (receiver *ImMemoryURL) GetURL(ctx context.Context) error {
	URLs, err := receiver.repo.GetURL(context.Background())
	if err != nil {
		return fmt.Errorf("Error while loading URL: %v", err)
	}
	fmt.Println(URLs)
	return nil
}
