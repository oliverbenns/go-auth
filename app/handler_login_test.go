package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginGetHandler_NoUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	server := Server{}

	handler := server.LoginHandler()
	handler(w, r)

	if w.Code != http.StatusOK {
		t.Error("Does not respond with correct status")
	}
}

func TestLoginGetHandler_WithUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&testAuthCookie)
	server := Server{}
	server.env.JwtSecretKey = "12345"

	handler := server.LoginHandler()
	handler(w, r)

	if w.Code != http.StatusFound {
		t.Error("Does not redirect")
	}
}

func TestLoginPostHandler_NoUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := CreateAuthFormPostRequest("T3eDol69ctknCGyPCRHd@4161hUm0Sb.com", "qPWKXkkwjB")
	server := CreateServer()
	handler := server.LoginHandler()
	handler(w, r)

	body := w.Body.String()

	if !strings.Contains(body, "Invalid credentials") {
		t.Error("Does not present error message to user")
	}

	if w.Code != http.StatusUnauthorized {
		t.Error("Returns incorrect status")
	}
}

func TestLoginPostHandler_Invalid(t *testing.T) {
	w := httptest.NewRecorder()
	r := CreateAuthFormPostRequest("john@example.com", "wrong_password")
	server := CreateServer()

	handler := server.LoginHandler()
	handler(w, r)

	body := w.Body.String()

	if !strings.Contains(body, "Invalid credentials") {
		t.Error("Does not present error message to user")
	}

	if w.Code != http.StatusUnauthorized {
		t.Error("Returns incorrect status")
	}
}

func TestLoginPostHandler_Valid(t *testing.T) {
	w := httptest.NewRecorder()
	r := CreateAuthFormPostRequest("john@example.com", "123456")
	server := CreateServer()

	handler := server.LoginHandler()
	handler(w, r)

	header := w.Result().Header["Set-Cookie"][0]

	if !strings.Contains(header, "user_token=") {
		t.Error("Cookie value is not set")
	}

	if w.Code != http.StatusFound {
		t.Error("Does not redirect")
	}
}
