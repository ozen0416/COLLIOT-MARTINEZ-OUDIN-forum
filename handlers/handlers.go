package handlers

import (
	"crypto/rand"
	"forum/handlers/connection"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	_ "github.com/gorilla/sessions"
)

var (
	key   = make([]byte, 64)
	_, _  = rand.Read(key)
	store = sessions.NewCookieStore(key)
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
}

func HandleIndex(files []string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "cookie-forum-ynov")

		if request.URL.Path != "/" {
			HandleNotFound(files, writer, request)
			return
		}
		f := append(files, "templates/landing-page.html")
		tmpl := template.Must(template.ParseFiles(f...))

		// If not authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			tmpl.Execute(writer, nil)
			return
		}
		// If authenticated
		id := session.Values["id-user"].(int)
		tmpl.Execute(writer, struct {
			Id   int
			Mess string
		}{Id: id, Mess: "authentifié et mon id est : "})
	})
}

func HandleNotFound(files []string, writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	f := append(files, "templates/not-found.html")
	tmpl := template.Must(template.ParseFiles(f...))
	tmpl.Execute(writer, nil)
}

func HandleDashboard(files []string) {
	http.HandleFunc("/dashboard", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "cookie-forum-ynov")

		if request.URL.Path != "/dashboard" {
			HandleNotFound(files, writer, request)
		}

		//if not logged
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(writer, request, "/login", 302)
		}

		f := append(files, "templates/dua.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, nil)
		return
	})
}
