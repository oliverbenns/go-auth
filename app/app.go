package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/oliverbenns/go-auth/pg"
	"net/http"
	"os"
)

func main() {
	server := Server{}
	server.db = pg.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/logout", server.LogoutHandler())
	mux.HandleFunc("/login", server.LoginHandler())
	mux.HandleFunc("/sign-up", server.SignUpHandler())
	mux.HandleFunc("/", server.IndexHandler())

	port := os.Getenv("PORT")

	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		panic(err)
	}
}
