package main

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/repositories/stores"
	"context"
	"fmt"
)

/*
	Microservicio de Telegram (version 3) -> By Collete

	[ Crea un programa que haga:										]
	[	1- una consulta mediante el metodo get o post a una página web,	]
	[	2- guarde su estada, codigo de respuesta o error en un mensaje	]
	[	3- envie este estado mediante el bot de telegram				]

	[v2] -> En esta segunda version, implementamos el mismo codigo con objetos y metodos.
		- Definir struct para el bot de telegram y el cliente HTTP
		- Usar metodos para enviar los mensajes y hacer las solicitudes
		- - Usar variables de entorno para las constantes
		- Usar interfaces (conjunto de metodos) para los tipos para desacoplar la implementación de sú úso

	[v3] -> En esta tercera version aplicamos la architectura hexagonal para un
	mejor codigo, más escalable y flexible.
	- El programa ha de ser capaz de trabajar con un archivo Json par obtener URLs
	y ha de escribirlas en otro, ademas de su contenido y codigo de respuesta.
	- También lo vamos hacer con in-memory (memoria local) para teneer dos tipos de repositorios.

	Architectura Hexagonal:
	1. Diseño del puerto.
	2. Adaptacioón de repos(in-memory y json).
	3.

Cosas por hacer:
  - Ser mas especifico con los errores para un mejor debug
  - Usar context.Context (para manejar cancelaciones y tiempos limite)
  - El codigo es secuencial,se pueden usar gorutinas para manejar varias solicitudes simultaneas
  - Ejemplo de Mock para pruebas
*/
func main() {
	//TODO
	cfg := config.GetConfig()

	//Getting URLs from ImMemory
	fmt.Printf("Getting URLs from ImMemory:\n")
	ImUserstoreOne, err := stores.NewImUserStore(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	err = ImUserstoreOne.LoadURLs(context.Background())
	if err != nil {
		panic(err)
	}

	repoURLs, err := ImUserstoreOne.GetURLs(context.Background())
	for _, u := range repoURLs.URLs {
		fmt.Printf("the %d URL is: %v\n", u.Id, u.Url)
	}

	fmt.Printf("\nGetting URLs from JSON:\n")
	JsonStores := stores.NewJsonStores(context.Background(), cfg)
	JsonStores.LoadURLs(context.Background())
	repoURLs2, err := JsonStores.GetURLs(context.Background())
	for _, u := range repoURLs2.URLs {
		fmt.Printf("the %d URL is: %v\n", u.Id, u.Url)
	}
}
