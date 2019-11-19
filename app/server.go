package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/oliverbenns/go-auth/pg"
	"os"
)

type Env struct {
	jwtSecretKey string
}

type Server struct {
	db  *sql.DB
	env Env
}

func NewServer() Server {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return Server{
		db: pg.Init(),
		env: Env{
			jwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
	}
}
