package main

import (
	"context"
	"fmt"
	"log"
	"microServBack/config"
	"microServBack/internal/domain/services"
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

	//Checker is an instance to check status and store its response
	Checker := services.NewCheckerAdapt(context.Background(), ImStore)
	err = Checker.GetURLsStatus(context.Background())
	if err != nil {
		log.Fatalf("error checking URLs from main: %v\n", err)
	}
	//Check if status have been stored correctly
	for i, u := range ImStore.UrlsDataResp {
		fmt.Printf("URL #%d: %v\n", i, u)
	}

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
