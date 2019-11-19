package main

import (
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Id    int64
	Email string
}

func main() {
	server := NewServer()
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
