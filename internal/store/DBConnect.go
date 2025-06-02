package store

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3333)/tasks_tracker")
	if err != nil {
		return nil, err
	}

	log.Println("[INFO] Database Connection was established successfully")

	return db, nil
}
