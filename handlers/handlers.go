package handlers

import (
	"html/template"
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

//func HandleLogin(files []string) {
//	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
//		if request.URL.Path != "login" {
//			HandleNotFound(files, writer, request)
//			return
//		}
//		if request.Method != http.MethodPost {
//			f := append(files, "templates/login.html")
//			tmpl := template.Must(template.ParseFiles(f...))
//			tmpl.Execute(w, nil)
//			return
//		}
//
//	})
//
//}
