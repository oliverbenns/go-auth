package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("views/layout.tmpl", "views/login.tmpl"))

type Alert struct {
	Message string
	Theme   string
}

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

func isAuthenticatedUser(r *http.Request) bool {
	cookie, err := r.Cookie("user_token")

	if err != nil {
		return false
	}

	// @TODO: validate jwt.
	validUser := cookie.Value == "abc"

	return validUser
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if isAuthenticatedUser(r) {
			http.Redirect(w, r, "/account", http.StatusFound)
		} else {
			tmpl.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		// email := r.FormValue("email")
		password := r.FormValue("password")
		// r.
		// @TODO: lookup user in DB
		// qwerty
		hash := "$2a$10$l3Lm6n/GIm9.j8/DTe05seV8E/uUPsh3Ie4NK08ncVUxLKRCnFqcK"
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		validCredentials := err == nil

		if validCredentials {
			// @TODO issue jwt
			jwt := "abc"
			setUserToken(w, jwt)
			alert := Alert{"Success!", "success"}
			tmpl.Execute(w, alert)
			// @TODO: Redirect
			// http.Redirect(w, r, "/account", http.StatusFound)
		} else {
			alert := Alert{"Invalid credentials. Please try again.", "danger"}
			tmpl.Execute(w, alert)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
