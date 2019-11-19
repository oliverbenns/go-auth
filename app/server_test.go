package main

import (
	"github.com/oliverbenns/go-auth/env"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()

	eenv := env.GetEnv()

	if server.env.Port != eenv.Port || server.env.Port == 0 {
		t.Error("Server env failed to load")
	}

	err := server.db.Ping()

	if err != nil {
		t.Error("Postgres failed to connect")
	}
}
