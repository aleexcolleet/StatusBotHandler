package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

/*
 */
type Config struct {
	//Bot      notifier.Bot
	//Chat     notifier.Chats
	JsonFiles JsonFiles
	//Postgres Postgres
}

/*
Postgres if I happened to use it, is an actual repo

type Postgres struct {
}
*/

/*
JsonFiles is a struct with the path to:
 1. Json file to acquire the Url to request
 2. Json file (which acts as a repo) to store Json URLS and requests
*/
type JsonFiles struct {
	JsonURLRoute  string
	JsonRepoRoute string
}

func GetConfig() Config {
	LoadEnvConfig()

	//botToken := os.Getenv("BOT_TOKEN")
	//desiredUrl := os.Getenv("Desired_URL")
	//chatIdStr := os.Getenv("CHAT_ID")
	JsonURLRoute := os.Getenv("JSON_URL_ROUTE")
	JsonRepoRoute := os.Getenv("JSON_REPO_ROUTE")

	if JsonURLRoute == "" || JsonRepoRoute == "" {
		fmt.Errorf(".env is empty\n")
	}
	/*
		if botToken == "" || desiredUrl == "" || chatIdStr == "" || JsonURLRoute == "" || JsonRepoRoute == "" {
			panic(fmt.Errorf("there's an empty variable in .env file"))
		}
		botLoad := notifier.Bot{
		ApiToken: botToken,
		}
		//botLoad.SetUrlApi()
		//We take different chatIds. Simply split it and store it in a slice
		chatLoad := notifier.Chats{}
		chatLoad.ChatId = strings.Split(chatIdStr, ",")
	*/
	jsonFiles := JsonFiles{
		JsonURLRoute:  JsonURLRoute,
		JsonRepoRoute: JsonRepoRoute,
	}

	return Config{
		//Bot:       botLoad,
		//Chat:      chatLoad,
		JsonFiles: jsonFiles,
	}
}

// LoadEnvConfig is a func to load specifically from ".env"
func LoadEnvConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %s", err))
	}
}
