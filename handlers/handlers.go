package handlers

import (
	"crypto/rand"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	_ "github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var (
	key   = make([]byte, 64)
	_, _  = rand.Read(key)
	store = sessions.NewCookieStore(key)
)

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
		}{Id: id, Mess: "authentifié est mon id est : "})
	})
}

func HandleNotFound(files []string, writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	f := append(files, "templates/not-found.html")
	tmpl := template.Must(template.ParseFiles(f...))
	tmpl.Execute(writer, nil)
}

func HandleLogin(files []string) {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "cookie-forum-ynov")
		if request.URL.Path != "/login" {
			HandleNotFound(files, writer, request)
		}

		if request.Method != http.MethodPost {
			f := append(files, "templates/login.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, nil)
			return
		}

		// If authenticated
		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			http.Redirect(writer, request, "/dashboard", 302)
		}

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		email := request.FormValue("email")
		pass := request.FormValue("pass")

		var idUser int

		err = db.QueryRow("SELECT id FROM `users` WHERE users.email = ? AND users.passwd = ?", email, pass).Scan(&idUser)
		if err != nil {
			f := append(files, "templates/login.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, "Erreur dans l'e-mail ou le mot de passe")
			return
		}

		session.Values["authenticated"] = true
		session.Values["id-user"] = idUser
		session.Save(request, writer)
		http.Redirect(writer, request, "/dashboard", 302)

	})

}

func HandleDashboard(files []string) {
	http.HandleFunc("/dashboard", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "cookie-forum-ynov")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(writer, request, "/login", 302)
		}

		f := append(files, "templates/dua.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, nil)
		return
	})
}

func HandlerLogout() {
	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		store.MaxAge(-1)
		http.Redirect(writer, request, "/", 302)
	})
}
