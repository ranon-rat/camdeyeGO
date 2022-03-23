package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func openDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", "./database.db")
	return db

}
