package main

import (
	"fmt"
	"net/http"
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

	fmt.Println("Server started on port", server.env.Port)
	address := fmt.Sprintf(":%d", server.env.Port)
	err := http.ListenAndServe(address, mux)

	if err != nil {
		panic(err)
	}
}
