package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUpPostHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.AddCookie(&testAuthCookie)
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
