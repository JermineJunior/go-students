package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		return nil, err
	}
	return db, nil
}
