package main

import "testing"

func TestCreateUserToken(t *testing.T) {
	user := User{
		Id:    2,
		Email: "example@example.com",
	}

	got := CreateUserToken(user)
	expected := "test"

	if got != expected {
		t.Error()
	}
}
