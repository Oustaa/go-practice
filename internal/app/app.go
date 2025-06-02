package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/oustaa/go-practice/internal/store"
)

type Application struct {
	DB     *sql.DB
	Logger *log.Logger
}

func NewApplication() *Application {
	DB, err := store.Open()
	if err != nil {
		log.Printf("[Error] store.Open(): %v", err)
	}

	Logger := log.New(os.Stdin, "", log.Ldate|log.Ltime)

	app := &Application{
		DB,
		Logger,
	}

	return app
}
