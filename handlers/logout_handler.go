package handlers

import (
	"github.com/oliverbenns/go-auth/auth"
	"net/http"
)

func logoutPostHandler(w http.ResponseWriter, r *http.Request) {
	auth.SetUserToken(w, "")
	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		logoutPostHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
