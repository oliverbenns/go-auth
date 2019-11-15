package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/oliverbenns/go-auth/db"
	"io/ioutil"
)

func main() {
	db.InitDb()
	query, err := ioutil.ReadFile("setup.sql")

	if err != nil {
		panic(err)
	}

	_, err = db.Db.Exec(string(query))

	if err != nil {
		panic(err)
	}

	db.Db.Close()
}
