package main

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var loginTmpl = LoadTemplate("login.html")

var invalidCredentialsAlert = Alert{"Invalid credentials. Please try again.", "danger"}

func loginGetHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	user := GetUserToken(r, s.env.JwtSecretKey)

	if user != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		loginTmpl.Execute(w, nil)
	}
}

func loginPostHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	user := User{}
	user.Email = r.FormValue("email")
	password := r.FormValue("password")
	var hash string

	row := s.db.QueryRow("SELECT id, hash FROM users WHERE email=$1 LIMIT 1", user.Email)
	err := row.Scan(&user.Id, &hash)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			loginTmpl.Execute(w, invalidCredentialsAlert)
			return
		} else {
			panic(err)
		}
	}

	validCredentials := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil

	if validCredentials {
		token := CreateUserToken(user, s.env.JwtSecretKey)
		SetUserToken(w, token)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		loginTmpl.Execute(w, invalidCredentialsAlert)
	}
}

func (s *Server) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			loginGetHandler(w, r, s)
		case http.MethodPost:
			loginPostHandler(w, r, s)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
