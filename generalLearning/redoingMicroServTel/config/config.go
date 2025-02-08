package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"redoingMicroServTel/internal/notifier"
	"strings"
)

/*
	In config I return configurations from .env
*/

type Config struct {
	Bot            notifier.Bot
	Chats          notifier.Chats
	JsonURLRoute   string
	JsonStoreRoute string
}

func GetConfig() (Config, error) {
	//Acquire the config file
	err := loadEnv(context.Background())
	if err != nil {
		return Config{}, fmt.Errorf("error loading environment variables: %w", err)
	}

	// Get Bot Api to the struct and setApiURL
	botLoad := notifier.Bot{
		ApiURL: os.Getenv("API_TOKEN"),
	}
	botLoad.SetApiURL()
	//Fetch ChatsID
	chatLoad := notifier.Chats{}
	chatLoad.ChatsId = strings.Split(os.Getenv("CHAT_ID"), ",")
	//Fetch JsonURLRoutes and return a copy by value of the struct
	return Config{
		Bot:            botLoad,
		Chats:          chatLoad,
		JsonURLRoute:   os.Getenv("JSON_URL_ROUTE"),
		JsonStoreRoute: os.Getenv("JSON_REPO_ROUTE"),
	}, nil
}

func loadEnv(ctx context.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading .env file: %v\n", err.Error())
	}
	return nil
}
