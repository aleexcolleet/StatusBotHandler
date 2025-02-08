package main

import (
	"fmt"
	"redoingMicroServTel/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("cfg: %v\n", cfg)
}
