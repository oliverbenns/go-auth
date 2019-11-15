package handlers

import (
	"database/sql"
	"fmt"
	"github.com/oliverbenns/go-auth/auth"
	"github.com/oliverbenns/go-auth/db"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var loginTmpl = template.Must(template.ParseFiles("views/layout.tmpl", "views/login.tmpl"))

type Alert struct {
	Message string
	Theme   string
}

var invalidCredentialsAlert = Alert{"Invalid credentials. Please try again.", "danger"}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_token")
	validUser := err == nil && auth.ValidateToken(cookie.Value)

	if validUser {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		loginTmpl.Execute(w, nil)
	}
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	var hash string

	query := fmt.Sprintf("SELECT hash FROM users WHERE email='%s' LIMIT 1", email)
	row := db.Db.QueryRow(query)
	err := row.Scan(&hash)

	if err != nil {
		if err == sql.ErrNoRows {
			loginTmpl.Execute(w, invalidCredentialsAlert)
			return
		} else {
			panic(err)
		}
	}

	validCredentials := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil

	if validCredentials {
		user := auth.User{email}
		token := auth.CreateToken(user)
		auth.SetUserToken(w, token)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		loginTmpl.Execute(w, invalidCredentialsAlert)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginGetHandler(w, r)

	} else if r.Method == "POST" {
		loginPostHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
