package main

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

var signUpTmpl = LoadTemplate("sign_up.html")

func signUpGetHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	user := GetUserToken(r, s.env.JwtSecretKey)

	if user != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		signUpTmpl.Execute(w, nil)
	}
}

func signUpPostHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	user := User{}
	user.Email = r.FormValue("email")
	password := r.FormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	row := s.db.QueryRow("INSERT INTO users(email, hash) VALUES ($1, $2) RETURNING id", user.Email, hash)
	err = row.Scan(&user.Id)

	if err != nil {
		var message string

		// @TODO: Find a better way to do this.
		if strings.Contains(err.Error(), "unique constraint \"users_email_key\"") {
			message = "User already exists with that email."
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			message = "An unknown error occured."
			w.WriteHeader(http.StatusInternalServerError)
		}

		alert := Alert{message, "danger"}
		signUpTmpl.Execute(w, alert)
	} else {
		token := CreateUserToken(user, s.env.JwtSecretKey)
		SetUserToken(w, token)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (s *Server) SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			signUpGetHandler(w, r, s)

		} else if r.Method == http.MethodPost {
			signUpPostHandler(w, r, s)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
