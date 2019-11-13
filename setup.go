package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	query, err := ioutil.ReadFile("setup.sql")

	if err != nil {
		panic(err)
	}

	_, qerr := db.Exec(string(query))

	if qerr != nil {
		panic(qerr)
	}

	db.Close()
}
