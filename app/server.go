package main

import (
	"database/sql"
	"github.com/oliverbenns/go-auth/env"
	"github.com/oliverbenns/go-auth/pg"
)

type Server struct {
	db  *sql.DB
	env env.Env
}

func NewServer() Server {
	serverEnv := env.GetEnv()

	return Server{
		db:  pg.Init(serverEnv),
		env: serverEnv,
	}
}
