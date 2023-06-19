package handlers

import (
	"fmt"
	"forum/handlers/connection"
	not_found "forum/handlers/not-found"
	"forum/session"
	"forum/structures"
	"html/template"
	"net/http"
	"strconv"
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
	connection.DeleteAcc()
	HandleTopic(f)
	HandleMention(f)
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
		_session, _ := session.Get(request)

		var data structures.Error

		if request.URL.Path != "/topic" {
			not_found.HandleNotFound(files, writer, request)
		}

		idTopic := strings.TrimPrefix(request.URL.RequestURI(), "/topic?id=")
		// True if is on a specific topic
		if idTopic != "/topic" {
			if request.Method == http.MethodPost {
				idTopicConverted, _ := strconv.Atoi(idTopic)
				// ok is false if user is not logged
				if idUser, ok := _session.Values["idUser"].(int); ok {
					mess := structures.Message{
						Content: request.FormValue("mess-content"),
						Author: structures.User{
							Id: idUser,
						},
						Topic: structures.Topic{
							Id: idTopicConverted,
						},
					}
					_, err := structures.SendMess(mess)
					if err != nil {
						fmt.Println("Send mess: ", err)
					}
				} else {
					data.ErrorMess = "Vous n'etes pas connecté"
				}
			}
			data.Data = structures.GetMessByTopicId(idTopic)
			f := append(files, "templates/topic-unique-temp.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(writer, data)
			return
		}
		// Executes if user is on the /topic route
		if request.Method == http.MethodPost {
			if idUser, ok := _session.Values["idUser"].(int); ok {
				topic := structures.Topic{
					Content: request.FormValue("topic-content"),
					CatId:   1,
					Author: structures.User{
						Id: idUser,
					},
				}

				err := structures.SendTopic(topic)
				if err != nil {
					fmt.Println("Send topic: ", err)
				}
			} else {
				data.ErrorMess = "Vous n'etes pas connecté"
			}
		}

		data.Data = structures.GetTopicsByTime()
		f := append(files, "templates/topics.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, data)
		return
	})
}

func HandleMention(files []string) {
	http.HandleFunc("/mentions", func(writer http.ResponseWriter, request *http.Request) {
		f := append(files, "templates/mentions.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(writer, nil)
	})
}
