package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJpZCI6Mn0.fsA-0yhLc_XwndToIxmytRkBmvD78akk1mkJ7Be_xNs"

	cookie := http.Cookie{
		Name:  "user_token",
		Value: token,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&cookie)
	server := Server{}
	server.env.JwtSecretKey = "12345"

	handler := server.LoginHandler()
	handler(w, r)

	if w.Code != http.StatusFound {
		t.Error("Does not redirect")
	}
}

func createPostRequest(email string, password string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)
	r.URL.RawQuery = form.Encode()

	return r
}

func createServer() *Server {
	// @TODO: This now uses the env.
	// Do we want a seperate test env loaded including a dummy secret key?
	server := NewServer()

	return &server
}

func TestLoginPostHandler_NoUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := createPostRequest("T3eDol69ctknCGyPCRHd@4161hUm0Sb.com", "qPWKXkkwjB")
	server := createServer()
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
	r := createPostRequest("john@example.com", "wrong_password")
	server := createServer()

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
	r := createPostRequest("john@example.com", "123456")
	server := createServer()

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
