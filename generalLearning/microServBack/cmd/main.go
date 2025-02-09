package main

import (
	"context"
	"microServBack/config"
)

func main() {
	cfg, err := config.GetConfig(context.Background())
	if err != nil {
		panic("config file error")
	}

	//For unused warning
	_ = cfg
}
