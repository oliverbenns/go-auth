package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"os"
)

var tmpl = template.Must(template.ParseFiles("views/layout.tmpl", "views/login.tmpl"))

type Alert struct {
	Message string
	Theme   string
}

type User struct {
	Email string
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

	validUser := validateToken(cookie.Value)

	fmt.Println("validUser", validUser)

	return validUser
}

func createToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	secret := []byte(secretKey)

	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func validateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		secret := []byte(secretKey)

		return secret, nil
	})

	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims["email"] == "abc@abc.com"
	} else {
		return false
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if isAuthenticatedUser(r) {
		http.Redirect(w, r, "/account", http.StatusFound)
	} else {
		tmpl.Execute(w, nil)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	// r.
	// @TODO: lookup user in DB
	// qwerty
	hash := "$2a$10$l3Lm6n/GIm9.j8/DTe05seV8E/uUPsh3Ie4NK08ncVUxLKRCnFqcK"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	validCredentials := err == nil

	if validCredentials {
		user := User{email}
		token := createToken(user)
		setUserToken(w, token)
		alert := Alert{"Success!", "success"}
		tmpl.Execute(w, alert)
		// @TODO: Redirect
		// http.Redirect(w, r, "/account", http.StatusFound)
	} else {
		alert := Alert{"Invalid credentials. Please try again.", "danger"}
		tmpl.Execute(w, alert)
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
