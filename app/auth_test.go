package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserToken(t *testing.T) {
	user := User{
		Id:    2,
		Email: "example@example.com",
	}

	got := CreateUserToken(user, "12345")
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJpZCI6Mn0.fsA-0yhLc_XwndToIxmytRkBmvD78akk1mkJ7Be_xNs"
	if got != expected {
		t.Error("Incorrect token generated")
	}
}

func TestParseUserToken(t *testing.T) {
	user := User{
		Id:    2,
		Email: "example@example.com",
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJpZCI6Mn0.fsA-0yhLc_XwndToIxmytRkBmvD78akk1mkJ7Be_xNs"
	userClaims, _ := ParseUserToken(token, "12345")

	if userClaims.Id != user.Id || userClaims.Email != user.Email {
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
	user := User{
		Id:    2,
		Email: "example@example.com",
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJpZCI6Mn0.fsA-0yhLc_XwndToIxmytRkBmvD78akk1mkJ7Be_xNs"

	cookie := http.Cookie{
		Name:  "user_token",
		Value: token,
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&cookie)
	userFromToken := GetUserToken(r, "12345")

	if userFromToken.Id != user.Id || userFromToken.Email != user.Email {
		t.Error("Unable to get user from cookie")
	}
}
