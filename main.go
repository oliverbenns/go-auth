package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/account", accountHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/sign-up", signUpHandler)
	mux.HandleFunc("/", indexHandler)
	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, mux)
}
