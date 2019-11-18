package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/oliverbenns/go-auth/pg"
	"io/ioutil"
)

func main() {
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
