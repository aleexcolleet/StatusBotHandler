package config

import (
	"context"
	"github.com/joho/godotenv"
	"microServBack/internal/domain/notifier"
	"os"
	"strings"
)

type Config struct {
	Bot           notifier.Bot
	Chat          notifier.Chat
	JsonUrlRoute  string
	JsonRepoRoute string
}

func GetConfig(ctx context.Context) (Config, error) {

	loadEnv()

	botLoad := notifier.Bot{
		ApiToken: os.Getenv("API_TOKEN"),
	}
	botLoad.SetUrlAPI()
	chatLoad := notifier.Chat{
		ChatId: strings.Split(os.Getenv("CHAT_TOKEN"), ","),
	}
	return Config{
		Bot:           botLoad,
		Chat:          chatLoad,
		JsonUrlRoute:  os.Getenv("JSON_URL_ROUTE"),
		JsonRepoRoute: os.Getenv("JSON_REPO_ROUTE"),
	}, nil

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
