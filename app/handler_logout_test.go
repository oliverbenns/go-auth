package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUpPostHandler(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJpZCI6Mn0.fsA-0yhLc_XwndToIxmytRkBmvD78akk1mkJ7Be_xNs"

	cookie := http.Cookie{
		Name:  "user_token",
		Value: token,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.AddCookie(&cookie)
	server := Server{}
	server.env.JwtSecretKey = "12345"

	handler := server.LogoutHandler()
	handler(w, r)

	header := w.Result().Header["Set-Cookie"][0]

	if !strings.Contains(header, "user_token=;") {
		t.Error("Cookie value is not set to empty")
	}

	if w.Code != http.StatusFound {
		t.Error("Does not redirect")
	}
}
