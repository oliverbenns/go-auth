package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSignUpGetHandler_NoUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	server := Server{}

	handler := server.SignUpHandler()
	handler(w, r)

	if w.Code != http.StatusOK {
		t.Error("Does not respond with correct status")
	}
}

func TestSignUpGetHandler_WithUser(t *testing.T) {
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

	handler := server.SignUpHandler()
	handler(w, r)

	if w.Code != http.StatusFound {
		t.Error("Does not redirect")
	}
}

func TestSignUpPostHandler_ExistingUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := createPostRequest("john@example.com", "qPWKXkkwjB")
	server := createServer()
	handler := server.SignUpHandler()
	handler(w, r)

	body := w.Body.String()

	if !strings.Contains(body, "User already exists") {
		t.Error("Does not present error message to user")
	}

	if w.Code != http.StatusUnprocessableEntity {
		t.Error("Returns incorrect status")
	}
}

func TestSignUpPostHandler_Success(t *testing.T) {
	w := httptest.NewRecorder()
	now := time.Now()
	email := fmt.Sprintf("test-%d@example.com", now.Unix())
	password := string(now.Unix())
	r := createPostRequest(email, password)
	server := createServer()

	handler := server.SignUpHandler()
	handler(w, r)

	header := w.Result().Header["Set-Cookie"][0]

	if !strings.Contains(header, "user_token=") {
		t.Error("Cookie value is not set")
	}

	if w.Code != http.StatusFound {
		t.Error("Does not redirect")
	}
}
