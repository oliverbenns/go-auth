package main

import (
	"github.com/oliverbenns/go-auth/env"
	"github.com/oliverbenns/go-auth/pg"
	"io/ioutil"
)

func main() {
	eenv := env.GetEnv()

	db := pg.Init(eenv)
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
