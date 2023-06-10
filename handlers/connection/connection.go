package connection

import (
	"database/sql"
	"fmt"
	"forum/handlers/not-found"
	"forum/secret"
	"forum/session"
	"html/template"
	"log"
	"net/http"
)

func HandleLogin(files []string) {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		_session, _ := session.Get(request)
		//Check if not 404
		if request.URL.Path != "/login" {
			not_found.HandleNotFound(files, writer, request)
		}

		//Check if it's first time on page
		if request.Method != http.MethodPost {
			// If authenticated
			if auth, ok := _session.Values["authenticated"].(bool); ok && auth {
				http.Redirect(writer, request, "/dashboard", 302)
			}

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

		var hashPass string
		var idUser int
		err = db.QueryRow("SELECT passwd, id FROM `users` WHERE users.email = ?", email).Scan(&hashPass, &idUser)
		if err != nil {
			fmt.Println("Login: ", err)
		}
		//Check if password is good
		if !secret.CheckPasswordHash(pass, hashPass) {
			f := append(files, "templates/login.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, "Erreur dans l'e-mail ou le mot de passe !")
			return
		}

		fmt.Println("Logged in")
		_session.Values["authenticated"] = true
		_session.Values["id-user"] = idUser
		_session.Save(request, writer)
		http.Redirect(writer, request, "/dashboard", 302)
	})
}

func HandlerLogout() {
	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		_session, _ := session.Get(request)
		_session.Values["authenticated"] = "false"
		_session.Save(request, writer)
		session.MaxAge(-1)
		http.Redirect(writer, request, "/", 302)
	})
}

func HandlerSignIn(files []string) {
	http.HandleFunc("/signin", func(writer http.ResponseWriter, request *http.Request) {
		_session, _ := session.Get(request)

		if request.URL.Path != "/signin" {
			not_found.HandleNotFound(files, writer, request)
		}

		if request.Method != http.MethodPost {
			// If authenticated
			if auth, ok := _session.Values["authenticated"].(bool); ok && auth {
				http.Redirect(writer, request, "/dashboard", 302)
			}

			f := append(files, "templates/sign-in.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, nil)
			return
		}

		// Connecting to BD
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Fetching form values
		email := request.FormValue("email")
		pass := request.FormValue("pass")
		username := request.FormValue("username")
		date := request.FormValue("bdate")

		hash, err := secret.HashPassword(pass)
		if err != nil {
			fmt.Println(err)
		}

		_, err = db.Query("INSERT INTO `users`(`nickname`, `email`, `Date_Birth`, `passwd`) VALUES (?,?,?,?)", username, email, date, hash)
		if err != nil {
			fmt.Println(err)
			f := append(files, "templates/sign-in.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, "L'e-mail ou le pseudo sont déjà utilisés !")
			return
		}
		fmt.Println("User created")
		_session.Values["authenticated"] = true
		_session.Save(request, writer)
		http.Redirect(writer, request, "/dashboard", 302)
	})
}
