package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "asdasd@example.org",
		Password: "password",
	}
}
