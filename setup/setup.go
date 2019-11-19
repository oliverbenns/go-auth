package main

import (
	"github.com/joho/godotenv"
	"github.com/oliverbenns/go-auth/pg"
	"io/ioutil"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db := pg.Init()
	query, err := ioutil.ReadFile("setup/setup.sql")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(query))

	if err != nil {
		panic(err)
	}

	db.Close()
}
