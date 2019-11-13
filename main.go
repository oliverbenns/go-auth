package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
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
	mux := http.NewServeMux()

	mux.HandleFunc("/account", accountHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/sign-up", signUpHandler)
	mux.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")

	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, mux)
}
