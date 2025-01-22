package main

import "fmt"
import "net/http"

func main() {

	//creating the server connection with a struct
	//	adding specific addres(port) and handler func
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}
	//error handler
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server. err num: ", err)
	}
}

// handler funct to act when server is created
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))

}
