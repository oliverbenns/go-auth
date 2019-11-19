package main

import (
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
