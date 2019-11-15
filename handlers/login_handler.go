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

var tmpl = template.Must(template.ParseFiles("views/layout.tmpl", "views/login.tmpl"))

type Alert struct {
	Message string
	Theme   string
}

var invalidAlert = Alert{"Invalid credentials. Please try again.", "danger"}

func setUserToken(w http.ResponseWriter, userToken string) {
	cookie := http.Cookie{
		Name:     "user_token",
		Value:    userToken,
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_token")
	validUser := err == nil && auth.ValidateToken(cookie.Value)

	fmt.Println("validUser", validUser)

	if validUser {
		http.Redirect(w, r, "/account", http.StatusFound)
	} else {
		tmpl.Execute(w, nil)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	var hash string

	query := fmt.Sprintf("SELECT hash FROM users WHERE email='%s' LIMIT 1", email)
	row := db.Db.QueryRow(query)
	err := row.Scan(&hash)

	if err != nil {
		if err == sql.ErrNoRows {
			tmpl.Execute(w, invalidAlert)
			return
		} else {
			panic(err)
		}
	}

	validCredentials := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil

	if validCredentials {
		user := auth.User{email}
		token := auth.CreateToken(user)
		setUserToken(w, token)
		http.Redirect(w, r, "/account", http.StatusFound)
	} else {
		tmpl.Execute(w, invalidAlert)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getHandler(w, r)

	} else if r.Method == "POST" {
		postHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
