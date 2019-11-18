package main

import (
	"database/sql"
)

type Server struct {
	// db     *someDatabase
	// router *someRouter
	db *sql.DB
}
