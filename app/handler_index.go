package main

import (
	"net/http"
)

var indexTmpl = LoadTemplate("index.html")

func indexGetHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	user := GetUserToken(r, s.env.JwtSecretKey)

	if user != nil {
		indexTmpl.Execute(w, user)
	} else {
		indexTmpl.Execute(w, nil)
	}
}

func (s *Server) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			indexGetHandler(w, r, s)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
