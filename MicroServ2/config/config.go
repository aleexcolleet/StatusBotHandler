package config

import (
	"MicroServ2/internal/domain/notifier"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	Bot      notifier.Bot
	Chat     notifier.Chats
	JsonFile JsonFiles
}

type JsonFiles struct {
	JsonUrlRoute  string
	//confiJsonRepoRoute string
}

func GetConfig() (Config, error) {

	err := loadEnv()
	if err != nil {
		fmt.Println("Error loading env file")
		return Config{}, err
	}

	botLoad := notifier.Bot{
		APIToken: os.Getenv("API_TOKEN"),
	}
	botLoad.SetUrlApi()
	chatLoad := notifier.Chats{}
	chatLoad.ChatsId = strings.Split(os.Getenv("CHATS_ID"), ",")

	jsonFileLoad := JsonFiles{
		JsonUrlRoute:  os.Getenv("JSON_URL_ROUTE"),
		//JsonRepoRoute: os.Getenv("JSON_REPO_ROUTE"),
	}

	return Config{
		Bot:      botLoad,
		Chat:     chatLoad,
		JsonFile: jsonFileLoad,
	}, nil

}

func loadEnv() error {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Cannot load .env")
		return err
	}
	return nil
}
