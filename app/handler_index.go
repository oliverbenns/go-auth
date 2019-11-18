package main

import (
	"net/http"
)

var indexTmpl = LoadTemplate("index.html")

func indexGetHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	user := GetUserToken(r)

	if user != nil {
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
