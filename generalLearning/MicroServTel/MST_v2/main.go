package main

/*
Falta un .env para que funcione
	Microservicio de Telegram (version 1) -> By Collete

	[ Crea un programa que haga:										]
	[	1- una consulta mediante el metodo get o post a una página web,	]
	[	2- guarde su estada, codigo de respuesta o error en un mensaje	]
	[	3- envie este estado mediante el bot de telegram				]

	[v2] -> En esta segunda version, implementamos el mismo codigo con objetos y metodos.
		- Definir struct para el bot de telegram y el cliente HTTP
		- Usar metodos para enviar los mensajes y hacer las solicitudes
		- - Usar variables de entorno para las constantes
		- Usar interfaces (conjunto de metodos) para los tipos para desacoplar la implementación de sú úso
	Cosas por hacer:
		- Ser mas especifico con los errores para un mejor debug
		- (falta por investigar esto junto a la architectura hexagonal.
		- Usar context.Context (para manejar cancelaciones y tiempos limite)
		- El codigo es secuencial,se pueden usar gorutinas para manejar varias solicitudes simultaneas
		- Ejemplo de Mock para pruebas

*/
import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	tlgrmBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Structs
// TelegramBot representa un bot de Telegram creado con BotFather
type TelegramBot struct {
	bot *tlgrmBotApi.BotAPI
}

// HTTPClient representa un Cliente HTTP desde el que mandar solicitudes
type HTTPClient struct {
	client *http.Client
}

// Interfaces
// Messagenger is an interface for message sending with Telegram
type Messenger interface {
	SendTelegramMessage(chatID int64, message string) error
}

// Requester is an interface to request a URL by method (GET / POST)
type Requester interface {
	MakeHTTPRequest(method string, url string, payload []byte) (string, *http.Response, error)
}

// Metodos
// NewTelegramBot es un Constructor del Bot de Telegram.
//
//	[Inicializa una instancia nueva de un bot de telegram]
func NewTelegramBot(token string) (*TelegramBot, error) {
	tmpBot, err := tlgrmBotApi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating a telegram bot")
	}
	return &TelegramBot{bot: tmpBot}, nil
}

// Constructor Client HTTP
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{client: &http.Client{}}
}

// Enviar mensajes
func (tb *TelegramBot) SendTelegramMessage(chatID int64, message string) error {
	msg := tlgrmBotApi.NewMessage(chatID, message)
	_, err := tb.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message")
	}
	return nil
}

// Hacer solicitudes HTTP
func (hcl *HTTPClient) MakeHTTPRequest(method string, url string, payload []byte) (string, *http.Response, error) {

	//Configuramos la req con metodo, la url y 0(valores a mandar)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", nil, fmt.Errorf("error creating the http request: %v", err)
	}

	//Mandamos la req y retorna la instancia de la resp
	resp, err := hcl.client.Do(req)
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

/*
	Para las variables de entorno necesitamos la libreria: github.com/joho/godotenv
	Una vez añadida, creadas las variables de entorno en un .env y exportadas.
	Cargamos las variables en el Main. Se important como strings. Por eso hay
		que transformar a telegramChatID
*/

func main() {

	//Accedemos a las variables de entorno lo antes posible
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("error loading .env file")
	}
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	tmpTelegramChatID := (os.Getenv("TELEGRAM_CHAT_ID"))
	desiredURL := os.Getenv("DESIRED_URL")
	if tmpTelegramChatID == "" || telegramBotToken == "" || desiredURL == "" {
		fmt.Errorf("telegram bot token is empty")
	}
	telegramChatID, err := strconv.ParseInt(tmpTelegramChatID, 10, 64)
	if err != nil {
		fmt.Errorf("error converting string to int64")
	}

	//Creamos una instancia de HTTPClient que implementa Requester
	var requester Requester = NewHTTPClient()

	//Realizamos la solicitud HTTP (GET o POST)
	body, resp, err := requester.MakeHTTPRequest("GET", desiredURL, nil)
	if err != nil {
		log.Fatalf("ERROR making http request: %v", err)
	}

	//Aqui preparamos el mensaje. Respuesta del cliente y la propia respuesta
	message := fmt.Sprintf("Codigo de respuesta: %d\n\nRespuesta: %s", resp.StatusCode, body)

	// Configuramos el Bot de Telegram y le hacemos mandar el mensaje de respuesta
	var messenger Messenger
	messenger, err = NewTelegramBot(telegramBotToken)
	if err != nil {
		log.Fatalf("Error creating the bot: %v", err)
	}
	err = messenger.SendTelegramMessage(telegramChatID, message)
	if err != nil {
		log.Fatalf("Error enviando el mensaje: %v", err)
	}

	fmt.Println("Mensaje enviado correctamente\n")
}
