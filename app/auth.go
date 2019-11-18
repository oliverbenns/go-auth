package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

type User struct {
	Id    int64
	Email string
}

func (s *Server) CreateToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	secret := []byte(secretKey)

	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func (s *Server) ValidateToken(tokenString string) bool {
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

	if !ok || !token.Valid {
		return false
	}

	row := s.db.QueryRow(`SELECT FROM users WHERE id=$1 LIMIT 1`, claims["id"])
	err = row.Scan()

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			panic(err)
		}
	}

	return true
}

func (s *Server) SetUserToken(w http.ResponseWriter, userToken string) {
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
