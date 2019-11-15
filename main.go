package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/oliverbenns/go-auth/db"
	"github.com/oliverbenns/go-auth/handlers"
	"net/http"
	"os"
)

func main() {
	db.InitDb()

	rows, _ := db.Db.Query("SELECT * FROM users")
	defer rows.Close()

	for rows.Next() {
		var (
			id         int64
			email      string
			hash       string
			created_at string
		)
		if err := rows.Scan(&id, &email, &hash, &created_at); err != nil {
			panic(err)
		}
		fmt.Printf("id %d email is %s\n", id, email)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/account", handlers.AccountHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/sign-up", handlers.SignUpHandler)
	mux.HandleFunc("/", handlers.IndexHandler)

	port := os.Getenv("PORT")

	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, mux)
}
