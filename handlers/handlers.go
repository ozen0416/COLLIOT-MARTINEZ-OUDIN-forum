package handlers

import (
	"database/sql"
	"fmt"
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
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM `topic` ORDER BY topic.publi_date DESC")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/login" {
			HandleNotFound(files, writer, request)
			return
		}
		for results.Next() {
			var content int
			results.Scan(&content)
			fmt.Println(content)
		}

		f := append(files, "templates/dua.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, nil)

		//if request.Method != http.MethodPost {
		//	f := append(files, "templates/login.html")
		//	tmpl := template.Must(template.ParseFiles(f...))
		//	tmpl.Execute(w, nil)
		//	return
		//}

	})

}
