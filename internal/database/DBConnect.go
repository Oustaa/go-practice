package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConncet() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@/tasks_tracker")
	if err != nil {
		return nil, err
	}

	return db, nil
}
