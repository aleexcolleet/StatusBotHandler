package config

import (
	"errors"
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
	Server     ServerFiber
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

type ServerFiber struct {
	Host string
	Port string
}

func GetConfig() (Config, error) {

	loadEnv()
	//variables required validator
	requiredVars := []string{"API_TOKEN", "CHAT_IDS", "DB_PORT", "DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, key := range requiredVars {
		if os.Getenv(key) == "" {
			return Config{}, errors.New("environment variable " + key + " not set")
		}
	}

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

	serverFiberLoad := ServerFiber{
		Host: os.Getenv("SRV_HOST"),
		Port: os.Getenv("SRV_PORT"),
	}

	return Config{
		Bot:   botLoad,
		Chats: chatLoad,
		JsonRoutes: JsonRoutes{
			JsonRouteRepo:      os.Getenv("JSON_ROUTE_REPO"),
			JsonRouteUrlSource: os.Getenv("JSON_ROUTE_URL_SOURCE"),
		},
		Server:   serverFiberLoad,
		Database: dbLoad,
	}, nil
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
