// Package session is needed in order to use cookies in different packages
package session

import (
	"crypto/rand"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key   = make([]byte, 64)
	_, _  = rand.Read(key)
	store = sessions.NewCookieStore(key)
)

func Get(req *http.Request) (*sessions.Session, error) {
	store.MaxAge(30)
	return store.Get(req, "cookie-forum-ynov")
}

func MaxAge(age int) {
	store.MaxAge(age)
}
