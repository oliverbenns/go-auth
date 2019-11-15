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

	mux := http.NewServeMux()

	mux.HandleFunc("/account", handlers.AccountHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/sign-up", handlers.SignUpHandler)
	mux.HandleFunc("/", handlers.IndexHandler)

	port := os.Getenv("PORT")

	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		panic(err)
	}
}
