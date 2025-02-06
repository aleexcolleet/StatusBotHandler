package main

/*
	Falta por hacer:
		- Checker

*/
import (
	"MicroServ2/config"
	"MicroServ2/internal/domain/usecases"
	"MicroServ2/internal/repositories/stores"
	"context"
	"fmt"
)

func main() {
	//	cgf is an instance of config
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Errorf("error loading config: %v", err)
	}
	//	ImStore is an instance of database
	ImStore := stores.NewImStore()
	//
	Domain := usecases.NewDomainInMemory(context.Background(), ImStore)
	err = Domain.LoadURL(cfg)
	if err != nil {
		fmt.Errorf("error loading url: %v", err)
	}
	err = Domain.GetURL(context.Background())
	if err != nil {
		fmt.Errorf("error loading url: %v", err)
	}
}
