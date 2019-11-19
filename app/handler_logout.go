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
		switch r.Method {
		case http.MethodPost:
			logoutPostHandler(w, r, s)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
