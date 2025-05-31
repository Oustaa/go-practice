package api

import (
	"database/sql"
	"net/http"
)

type TasksHandler struct {
	DB *sql.DB
}

func (th *TasksHandler) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func GetTaskById(w http.ResponseWriter, r *http.Request) {

}
