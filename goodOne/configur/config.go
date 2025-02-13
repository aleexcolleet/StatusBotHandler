package configur

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	Bot       Bot
	Chats     Chats
	JsonFiles JsonFiles
}

type Bot struct {
	APIToken string
	ApiURL   string
}

type Chats struct {
	ChatsId []string
}
type JsonFiles struct {
	JsonUrlRoute  string
	JsonRepoRoute string
}

func GetConfig() (Config, error) {

	err := loadEnv()
	if err != nil {
		fmt.Println("Error loading .env file")
		return Config{}, err
	}

	botLoad := Bot{
		APIToken: os.Getenv("API_TOKEN"),
	}
	chatLoad := Chats{}
	chatLoad.ChatsId = strings.Split(os.Getenv("CHATS_ID"), ",")

	jsonFilesLoad := JsonFiles{
		JsonUrlRoute:  os.Getenv("JSON_URL_ROUTE"),
		JsonRepoRoute: os.Getenv("JSON_REPO_ROUTE"),
	}

	return Config{
		Bot:       botLoad,
		Chats:     chatLoad,
		JsonFiles: jsonFilesLoad,
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
