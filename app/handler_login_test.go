package main

import (
	"net/http"
	"net/http/httptest"
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
