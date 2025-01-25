package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	telegramBotToken = "7868583574:AAH2qTuLLRDtR2ruzxgK0y6cc_JZWwREcUU"
	telegramChatID   = int64(8115362810)
)

func sendTelegramMessage(bot *tlgrmBotApi.BotAPI, message string) error {

	//creamos el nuevo mensaje con su direccion
	msg := tlgrmBotApi.NewMessage(telegramChatID, message)

	//enviando el mensaje
	sendMessage, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending telegram message: %v", err)
	}
	fmt.Printf("el mensaje enviado a sido ---> %+v\n", sendMessage)
	return nil
}

func makeHTPPRequest(method string, url string, payload []byte) (string, *http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", nil, fmt.Errorf("error creating the http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error executing the http request: %v", err)
	}

	defer resp.Body.Close()

	//ahora almacenamos el cuerpo de la respuesta para mandarlo al telegram
	//en formato []bytes (array/slice de bytes) ya que NewRequest retorna un io.ReadAll
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("error reading the http response body: %v", err)
	}

	return string(bodyResp), resp, nil
}

func main() {

	//URL que vamos a consultar
	//desiredURL := "https://myWebCollete.com"
	desiredURL := "https://jsonplaceholder.typicode.com/posts/1"

	//Realizamos la solicitud HTTP (GET o POST)
	body, resp, err := makeHTPPRequest("GET", desiredURL, nil)
	//este tipo(log.Fatalf) imprime el error i termina el programa [os.Exit(1)]
	if err != nil {
		log.Fatalf("Error realizando la solicitud: %v", err)
	}

	//preparar el mensaje con el estado y la respuesta(cliente)
	message := fmt.Sprintf("Codigo de respuesta: %d\nRespuesta: %s", resp.StatusCode, body)

	//configurar el bot de telegram
	bot, err := tlgrmBotApi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatalf("Error creating the bot: %v", err)
	}
	//mandamos el mesaje al bot
	err = sendTelegramMessage(bot, message)
	if err != nil {
		log.Fatalf("Error enviando el mensaje: %v", err)
	} else {
		fmt.Println("Mensaje enviado correctamente\n")
	}
}
