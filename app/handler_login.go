package main

import (
	"database/sql"
	"fmt"
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
			loginTmpl.Execute(w, invalidCredentialsAlert)
			return
		} else {
			panic(err)
		}
	}

	validCredentials := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil

	fmt.Println(s.env)
	if validCredentials {
		token := CreateUserToken(user, s.env.JwtSecretKey)
		SetUserToken(w, token)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		loginTmpl.Execute(w, invalidCredentialsAlert)
	}
}

func (s *Server) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			loginGetHandler(w, r, s)

		} else if r.Method == "POST" {
			loginPostHandler(w, r, s)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
