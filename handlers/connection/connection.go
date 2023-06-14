package connection

import (
	"fmt"
	"forum/handlers/not-found"
	"forum/secret"
	"forum/session"
	"forum/structures"
	"html/template"
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

		email := request.FormValue("email")
		pass := request.FormValue("pass")

		//Check if password is good + get user id
		var idUser int
		var ok bool
		if idUser, ok = structures.Login(email, pass); ok != true {
			f := append(files, "templates/login.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, "Erreur dans l'e-mail ou le mot de passe !")
			return
		}

		fmt.Println("Logged in")
		_session.Values["authenticated"] = true
		fmt.Println(idUser)
		_session.Values["idUser"] = idUser
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

		// Fetching form values
		pass := request.FormValue("pass")
		hash, err := secret.HashPassword(pass)
		if err != nil {
			fmt.Println(err)
		}
		UserToSign := structures.User{
			Email:     request.FormValue("email"),
			Username:  request.FormValue("username"),
			Pass:      hash,
			BirthDate: request.FormValue("bdate"),
		}

		idUser, err := structures.SignUser(UserToSign)
		if err != nil {
			fmt.Println("Sign In: ", err)
			f := append(files, "templates/sign-in.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, "L'e-mail ou le pseudo sont déjà utilisés !")
			return
		}
		fmt.Println("User created")
		_session.Values["authenticated"] = true
		_session.Values["idUser"] = idUser
		_session.Save(request, writer)
		http.Redirect(writer, request, "/dashboard", 302)
	})
}

func DeleteAcc() {
	http.HandleFunc("/deleteacc", func(writer http.ResponseWriter, request *http.Request) {
		_session, _ := session.Get(request)
		_session.Values["authenticated"] = "false"
		idUser := _session.Values["idUser"]
		_session.Save(request, writer)
		structures.DeleteUser(idUser.(int))
		session.MaxAge(-1)
		http.Redirect(writer, request, "/", 302)
	})
}
