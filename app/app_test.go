package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var testUserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJpZCI6Mn0.fsA-0yhLc_XwndToIxmytRkBmvD78akk1mkJ7Be_xNs"

var testUser = User{
	Id:    2,
	Email: "example@example.com",
}

var testAuthCookie = http.Cookie{
	Name:  "user_token",
	Value: testUserToken,
}

func CreateAuthFormPostRequest(email string, password string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)
	r.URL.RawQuery = form.Encode()

	return r
}

func CreateServer() *Server {
	// @TODO: This now uses the env.
	// Do we want a seperate test env loaded including a dummy secret key?
	server := NewServer()

	return &server
}
