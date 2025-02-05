package main

import (
	"cmd/main.go/internal/repositories/stores"
	"context"
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
	1. Disenyo del puerto.

Cosas por hacer:
  - Ser mas especifico con los errores para un mejor debug
  - Usar context.Context (para manejar cancelaciones y tiempos limite)
  - El codigo es secuencial,se pueden usar gorutinas para manejar varias solicitudes simultaneas
  - Ejemplo de Mock para pruebas
*/
func main() {
	JsonStores := stores.NewJsonStores()
	JsonStores.LoadURLs(context.Background())
	
}
