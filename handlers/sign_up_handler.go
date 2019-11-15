package handlers

import (
	"fmt"
	"github.com/oliverbenns/go-auth/auth"
	"github.com/oliverbenns/go-auth/db"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strings"
)

var signUpTmpl = template.Must(template.ParseFiles("views/layout.html", "views/sign_up.html"))

func signUpGetHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_token")
	validUser := err == nil && auth.ValidateToken(cookie.Value)

	if validUser {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		signUpTmpl.Execute(w, nil)
	}
}

func signUpPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf("INSERT INTO users(email, hash) VALUES ('%s', '%s')", email, hash)
	_, err = db.Db.Exec(query)

	if err != nil {
		var message string

		// @TODO: Find a better way to do this.
		if strings.Contains(err.Error(), "unique constraint \"users_email_key\"") {
			message = "User already exists with that email."
		} else {
			message = "An unknown error occured."
		}

		alert := Alert{message, "danger"}
		signUpTmpl.Execute(w, alert)
	} else {
		user := auth.User{email}
		token := auth.CreateToken(user)
		auth.SetUserToken(w, token)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		signUpGetHandler(w, r)

	} else if r.Method == "POST" {
		signUpPostHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
