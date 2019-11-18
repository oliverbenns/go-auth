package main

import (
	"html/template"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("app/views/layout.html", "app/views/index.html"))

func indexGetHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	cookie, err := r.Cookie("user_token")
	validUser := err == nil && s.ValidateToken(cookie.Value)

	if validUser {
		user := User{"dummy@dummy.com"}
		indexTmpl.Execute(w, user)
	} else {
		indexTmpl.Execute(w, nil)
	}
}

func (s *Server) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			indexGetHandler(w, r, s)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
