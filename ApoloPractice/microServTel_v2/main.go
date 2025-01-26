package main

/*
	Microservicio de Telegram (version 1) -> By Collete

	[ Crea un programa que haga:										]
	[	1- una consulta mediante el metodo get o post a una página web,	]
	[	2- guarde su estada, codigo de respuesta o error en un mensaje	]
	[	3- envie este estado mediante el bot de telegram				]

	[v2] -> En esta segunda version, implementamos el mismo codigo con objetos y metodos.
		- Definir struct para el bot de telegram y el cliente HTTP
		- Usar metodos para enviar los mensajes y hacer las solicitudes

*/
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

// Structs
// Bot de Telegram
type TelegramBot struct {
	bot *tlgrmBotApi.BotAPI
}

// Cliente HTTP
type HTTPClient struct {
	client *http.Client
}

// Metodos
// 	Constructor del Bot de Telegram
// 		[Inicializa una instancia nueva de un bot de telegram]
func NewTelegramBot(token string) (*TelegramBot, error) {
	tmp_bot, err := tlgrmBotApi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating a telegram bot")
	}
	return &TelegramBot{bot: tmp_bot}, nil
}

// Constructor Client HTTP
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{client: &http.Client{}}
}

// Enviar mensajes
func (tb *TelegramBot) SendMessage(chatID int64, message string) error {
	msg := tlgrmBotApi.NewMessage(chatID, message)
	_, err := tb.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message")
	}
	return nil
}

func sendTelegramMessage(bot *tlgrmBotApi.BotAPI, message string) error {

	//creamos el nuevo mensaje con su destino (telegramChatId)
	msg := tlgrmBotApi.NewMessage(telegramChatID, message)

	//Enviamos el mensaje
	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending telegram message: %v", err)
	}
	return nil
}

func makeHTPPRequest(method string, url string, payload []byte) (string, *http.Response, error) {

	//Creamos la instancia del Cliente des de la que se hace la Req GET
	client := &http.Client{}

	//Configuramos la req con metodo, la url y 0(valores a mandar)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", nil, fmt.Errorf("error creating the http request: %v", err)
	}

	//Mandamos la req y retorna la instancia de la resp
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error executing the http request: %v", err)
	}

	//Dejamos en Defer la liberacion de la respuesta para que se ejecute al final
	defer resp.Body.Close()

	//Almacenamos el cuerpo de la respuesta para mandarlo al telegram
	//en formato []bytes (array/slice de bytes) ya que NewRequest retorna un io.ReadAll
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("error reading the http response body: %v", err)
	}

	return string(bodyResp), resp, nil
}

func main() {

	desiredURL := "https://jsonplaceholder.typicode.com/posts/1"

	//Realizamos la solicitud HTTP (GET o POST)
	body, resp, err := makeHTPPRequest("GET", desiredURL, nil)

	if err != nil {
		log.Fatalf("Error realizando la solicitud: %v", err)
	}

	//Aqui preparamos el mensaje. Respuesta del cliente y la propia respuesta
	message := fmt.Sprintf("Codigo de respuesta: %d\n\nRespuesta: %s", resp.StatusCode, body)

	//Configuramos el Bot de Telegram y le hacemos mandar el mensaje de respuesta
	bot, err := tlgrmBotApi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatalf("Error creating the bot: %v", err)
	}
	err = sendTelegramMessage(bot, message)
	if err != nil {
		log.Fatalf("Error enviando el mensaje: %v", err)
	}
	fmt.Println("Mensaje enviado correctamente\n")
}
