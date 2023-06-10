package handlers

import (
	"fmt"
	"forum/handlers/connection"
	not_found "forum/handlers/not-found"
	"forum/session"
	"forum/structures"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/sessions"
)

func OneHandlerToHandleThemAll() {
	f := []string{
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
	}

	HandleIndex(f)
	HandleDashboard(f)
	connection.HandleLogin(f)
	connection.HandlerLogout()
	connection.HandlerSignIn(f)
	HandleTopic(f)
}

func HandleIndex(files []string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_session, _ := session.Get(request)

		if request.URL.Path != "/" {
			not_found.HandleNotFound(files, writer, request)
			return
		}
		f := append(files, "templates/landing-page.html")
		tmpl := template.Must(template.ParseFiles(f...))

		// If not authenticated
		if auth, ok := _session.Values["authenticated"].(bool); !ok || !auth {
			tmpl.Execute(writer, nil)
			return
		}
		// If authenticated
		tmpl.Execute(writer, nil)
	})
}

func HandleDashboard(files []string) {
	http.HandleFunc("/dashboard", func(writer http.ResponseWriter, request *http.Request) {
		_session, _ := session.Get(request)

		if request.URL.Path != "/dashboard" {
			not_found.HandleNotFound(files, writer, request)
		}

		//if not logged
		if auth, ok := _session.Values["authenticated"].(bool); !ok || !auth {
			fmt.Println("Dashboard: auth: ", auth, " ok ?; ", ok)
			http.Redirect(writer, request, "/login", 302)
		}

		f := append(files, "templates/dua.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, nil)
		return
	})
}

func HandleTopic(files []string) {
	http.HandleFunc("/topic", func(writer http.ResponseWriter, request *http.Request) {
		//_session, _ := session.Get(request)

		if request.URL.Path != "/topic" {
			not_found.HandleNotFound(files, writer, request)
		}

		idTopic := strings.TrimPrefix(request.URL.RequestURI(), "/topic?id=")
		if idTopic != "/topic" {
			f := append(files, "templates/topic-unique-temp.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, structures.GetMessByTopicId(idTopic))
			return
		}
		f := append(files, "templates/topics.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, structures.GetTopicsByTime())
		return
	})
}
