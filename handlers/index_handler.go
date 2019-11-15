package handlers

import (
	"github.com/oliverbenns/go-auth/auth"
	"html/template"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("views/layout.html", "views/index.html"))

type User struct {
	Email string
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_token")
	validUser := err == nil && auth.ValidateToken(cookie.Value)

	if validUser {
		user := User{"dummy@dummy.com"}
		indexTmpl.Execute(w, user)
	} else {
		indexTmpl.Execute(w, nil)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		indexGetHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
