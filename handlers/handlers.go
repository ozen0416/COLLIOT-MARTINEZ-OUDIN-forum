package handlers

import (
	"database/sql"
	"forum/structures"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

func HandleIndex(files []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			HandleNotFound(files, w, r)
			return
		}
		f := append(files, "templates/landing-page.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, nil)
	})
}

func HandleNotFound(files []string, w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(404)
	f := append(files, "templates/not-found.html")
	tmpl := template.Must(template.ParseFiles(f...))
	tmpl.Execute(w, nil)
}

func HandleLogin(files []string) {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/login" {
			HandleNotFound(files, writer, request)
			return
		}

		if request.Method != http.MethodPost {
			f := append(files, "templates/login.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, nil)
			return
		}

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		email := request.FormValue("email")
		pass := request.FormValue("pass")

		var authValue structures.Auth

		err = db.QueryRow("SELECT email, passwd FROM `users` WHERE users.email = ? AND users.passwd = ?", email, pass).Scan(&authValue.Email, &authValue.Pass)
		if err != nil {
			f := append(files, "templates/login.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, "Erreur dans l'e-mail ou le mot de passe")
			return
		}

		http.Redirect(writer, request, "/dashboard", 302)

	})

}

func HandleDashboard(files []string) {
	http.HandleFunc("/dashboard", func(writer http.ResponseWriter, request *http.Request) {
		f := append(files, "templates/dua.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, nil)
		return
	})
}
