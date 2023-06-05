package main

import (
	"fmt"
	"forum/handlers"
	"net/http"
)

const port = ":8080"

func main() {

	f := []string{
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
	}

	// Ajout des fichiers statiques
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	handlers.HandleIndex(f)
	handlers.HandleLogin(f)
	http.ListenAndServe(port, nil)
}
