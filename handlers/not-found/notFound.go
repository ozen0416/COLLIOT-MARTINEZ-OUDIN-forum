package not_found

import (
	"html/template"
	"net/http"
)

func HandleNotFound(files []string, writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	f := append(files, "templates/not-found.html")
	tmpl := template.Must(template.ParseFiles(f...))
	tmpl.Execute(writer, nil)
}
