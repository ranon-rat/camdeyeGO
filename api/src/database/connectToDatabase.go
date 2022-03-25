package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	// env options
	dbEnv := os.Getenv("DB")
	dbHost := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	// uri database options
	uri := "postgres://postgres@127.0.0.1:5432/microblog?sslmode=disable"
	if dbEnv != "" {
		uri = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, dbHost, port, dbEnv)
	}
	// then i connect
	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err) // the api will stop
	}
	return db

}
