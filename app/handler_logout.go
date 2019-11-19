package main

import (
	"net/http"
)

func logoutPostHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	SetUserToken(w, "")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *Server) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			logoutPostHandler(w, r, s)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
