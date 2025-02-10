package main

import (
	"context"
	"microServBack/config"
	"microServBack/internal/domain/usecases"
	"microServBack/internal/repositories/stores"
)

func main() {
	cfg, err := config.GetConfig(context.Background())
	if err != nil {
		panic("config file error")
	}

	//Using loadUrls to load JsonFile URLs into the repo
	//Instance of the repository (database)
	ImStore := stores.NewImStore()
	Domain := usecases.NewImMemoryStore(context.Background(), ImStore)
	Domain.LoadUrl(cfg)

	/*
		I don't need to get Urls from here, but I'll leave the implementation below
		//GetUrls from domain accesses the repo adaptation and returns Urls according
		//to each repo.
		Urls, err := Domain.GetUrls(context.Background())
		if err != nil {
			log.Fatalf("error with urls getter in main: %v", err)
		}
	*/

}
