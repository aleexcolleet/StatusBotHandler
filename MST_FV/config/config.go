package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// Config is a package that I use to get config configuration.

type Config struct {
	Bot        Bot
	Chats      Chats
	JsonRoutes JsonRoutes
}
type Bot struct {
	ApiToken string
}

type Chats struct {
	ChatsId []string
}
type JsonRoutes struct {
	JsonRouteRepo      string
	JsonRouteUrlSource string
}

func GetConfig() (Config, error) {

	loadEnv()
	botLoad := Bot{
		ApiToken: os.Getenv("API_TOKEN"),
	}

	chatLoad := Chats{}
	chatLoad.ChatsId = strings.Split(os.Getenv("CHAT_IDS"), ",")

	return Config{
		Bot:   botLoad,
		Chats: chatLoad,
		JsonRoutes: JsonRoutes{
			JsonRouteRepo:      os.Getenv("JSON_ROUTE_REPO"),
			JsonRouteUrlSource: os.Getenv("JSON_ROUTE_URL_SOURCE"),
		},
	}, nil
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
