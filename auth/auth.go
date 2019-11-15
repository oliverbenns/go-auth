package auth

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/oliverbenns/go-auth/db"
	"os"
)

type User struct {
	Email string
}

func CreateToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	secret := []byte(secretKey)

	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func ValidateToken(tokenString string) bool {
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

	query := fmt.Sprintf("SELECT FROM users WHERE email='%s' LIMIT 1", claims["email"])
	row := db.Db.QueryRow(query)
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
