package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexGetHandler_NoUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	server := Server{}

	handler := server.IndexHandler()
	handler(w, r)

	body := w.Body.String()

	if strings.Contains(body, "You are logged in") {
		t.Error("Responds as if logged in")
	}
}

func TestIndexGetHandler_WithUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&testAuthCookie)
	server := Server{}
	server.env.JwtSecretKey = "12345"

	handler := server.IndexHandler()
	handler(w, r)

	body := w.Body.String()

	if strings.Contains(body, "You are not logged in") {
		t.Error("Responds as if not logged in")
	}
}
