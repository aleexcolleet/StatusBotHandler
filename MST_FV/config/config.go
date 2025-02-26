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
	Database   Database
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

type Database struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
}

func GetConfig() (Config, error) {

	loadEnv()
	botLoad := Bot{
		ApiToken: os.Getenv("API_TOKEN"),
	}

	chatLoad := Chats{}
	chatLoad.ChatsId = strings.Split(os.Getenv("CHAT_IDS"), ",")

	dbLoad := Database{
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return Config{
		Bot:   botLoad,
		Chats: chatLoad,
		JsonRoutes: JsonRoutes{
			JsonRouteRepo:      os.Getenv("JSON_ROUTE_REPO"),
			JsonRouteUrlSource: os.Getenv("JSON_ROUTE_URL_SOURCE"),
		},
		Database: dbLoad,
	}, nil
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
