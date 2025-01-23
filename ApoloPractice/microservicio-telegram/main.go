package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// esta funcion realiza una solicitud GET para retornar el estado y el cuerpo
func makeHTPPRequest(method string, url string, payload []byte) (string, *http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	//fmt.Errorf() -> lets us use formatting features to create error messages
	//no lo imprime por pantalla, solo lo acumula para manejarlo mas tarde
	if err != nil {
		return "", nil, fmt.Errorf("error creating the http request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error executing the http request: %v", err)
	}
	//esto pospone el cierre(liberacion de datos)
	defer resp.Body.Close()
	//ahora almacenamos el cuerpo de la respuesta para mandarlo al telegram
	//en formato []bytes (array/slice de bytes)
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("error reading the http response body: %v", err)
	}
	return string(bodyResp), resp, nil
}

func main() {

	//URL que vamos a consultar
	desiredURL := "https:myWebCollete.com"

	//Realizamos la solicitud HTTP (GET o POST)
	body, resp, err := makeHTPPRequest("GET", desiredURL, nil)
	//este tipo(log.Fatal) imprime el error i termina el programa [os.Exit(1)]
	if err != nil {
		log.Fatal("Error realizando la solicitud: %v", err)
	}

	//preparar el mensaje con el estado y la respuesta(cliente)
	message := fmt.Sprinf("Codigo de respuesta: %d\nRespuesta: %s", resp.StatusCode, body)

	//configurar el bot de telegram
}
