package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

type UserClaims struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateUserToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	secret := []byte(secretKey)

	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func ParseUserToken(tokenString string) (*UserClaims, error) {
	var token *jwt.Token
	var err error

	token, err = jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("JWT_SECRET_KEY")
		secret := []byte(secretKey)
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)

	if ok {
		return claims, nil
	} else {
		return nil, errors.New("Type assertion against user claims")
	}
}

func SetUserToken(w http.ResponseWriter, userToken string) {
	cookie := http.Cookie{
		Name:   "user_token",
		Value:  userToken,
		Domain: "localhost",
		Path:   "/",
		// MaxAge:   0,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func GetUserToken(r *http.Request) *User {
	cookie, err := r.Cookie("user_token")

	if err != nil {
		return nil
	}

	userClaims, err := ParseUserToken(cookie.Value)

	if err != nil {
		return nil
	}

	return &User{
		Id:    userClaims.Id,
		Email: userClaims.Email,
	}
}
