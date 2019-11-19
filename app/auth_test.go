package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserToken(t *testing.T) {
	token := CreateUserToken(testUser, "12345")

	if token != testUserToken {
		t.Error("Incorrect token generated")
	}
}

func TestParseUserToken(t *testing.T) {
	userClaims, _ := ParseUserToken(testUserToken, "12345")

	if userClaims.Id != testUser.Id || userClaims.Email != testUser.Email {
		t.Error("Parsed user is incorrect")
	}
}

func TestSetUserToken(t *testing.T) {
	w := httptest.NewRecorder()
	SetUserToken(w, "abc")

	header := w.Result().Header["Set-Cookie"]

	if len(header) == 0 {
		t.Error("Cookie header not set")
	}

	value := header[0]

	if value != "user_token=abc; Path=/; Domain=localhost; HttpOnly" {
		t.Error("Cookie header incorrectly set")
	}
}

func TestGetUserToken(t *testing.T) {
	cookie := http.Cookie{
		Name:  "user_token",
		Value: testUserToken,
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&cookie)
	userFromToken := GetUserToken(r, "12345")

	if userFromToken.Id != testUser.Id || userFromToken.Email != testUser.Email {
		t.Error("Unable to get user from cookie")
	}
}
