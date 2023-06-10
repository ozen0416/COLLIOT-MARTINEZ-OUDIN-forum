package main

import (
	"fmt"
	"forum/handlers"
	"net/http"
)

const port = ":8080"

func main() {

	// Ajout des fichiers statiques
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	handlers.OneHandlerToHandleThemAll()

	http.ListenAndServe(port, nil)
}
