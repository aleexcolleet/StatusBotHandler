package main

import (
	"fmt"
	"log"
	"net/http"

	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// esta funcion realiza una solicitud GET para retornar el estado y el cuerpo
func makeHTPPRequest(method string, url string, payload []byte) (string, *http.Response, error) {

}

func main() {

	//URL que vamos a consultar
	desiredURL := "https:myWebCollete.com"

	//Realizamos la solicitud HTTP (GET o POST)
	status, body, err := makeHTPPRequest("GET", desiredURL, nil)
	if err != nil {
		log.Fatal("Error realizando la solicitud: %v", err)
	}
}
