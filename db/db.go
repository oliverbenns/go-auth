package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var Db *sql.DB

func InitDb() {
	var err error

	port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	Db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = Db.Ping()

	if err != nil {
		panic(err)
	}
}
