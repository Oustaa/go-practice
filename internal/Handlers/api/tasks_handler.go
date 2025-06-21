package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-practice/internal/store"
)

type Taskshandler struct {
	TasksStore *store.MySQLTasksService
}

func NewTaskshandler(TasksStore *store.MySQLTasksService) *Taskshandler {
	return &Taskshandler{TasksStore}
}

func (th Taskshandler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.TasksStore.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if tasks == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (th Taskshandler) GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "id params is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "id params is invalid", http.StatusBadRequest)
		return
	}

	task, err := th.TasksStore.GetTaskById(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task.ToResponse())

}

func (th Taskshandler) PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Get The Data from the body
	var task store.Task // it is not the ideal way to get the Task struct
	json.NewDecoder(r.Body).Decode(&task)
	defer r.Body.Close()

	// Create the task
	th.TasksStore.CreateTask(task)

	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(task)
}

func (th Taskshandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "id params is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "id params is invalid", http.StatusBadRequest)
		return
	}

	err = th.TasksStore.DeleteTasks(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (th Taskshandler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "id params is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "id params is invalid", http.StatusBadRequest)
		return
	}

	task, err := th.TasksStore.GetTaskById(idInt)
	// handle the case where no tasks found, and if an error arise

	type TaskToUpdateInterface struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		CreatedAt   *string `json:"created_at"`
		UpdatedAt   *string `json:"updated_at"`
		CategoryID  *int    `json:"category_id"`
		StatusID    *int    `json:"status_id"`
	}

	bodyTask := &TaskToUpdateInterface{}
	json.NewDecoder(r.Body).Decode(&bodyTask)
	defer r.Body.Close()

	if bodyTask.Title != nil {
		task.Title = *bodyTask.Title
	}

	if bodyTask.Description != nil {
		task.Description = sql.NullString{String: *bodyTask.Description, Valid: true}
	}

	if bodyTask.CategoryID != nil {
		task.CategoryID = *bodyTask.CategoryID
	}

	if bodyTask.StatusID != nil {
		task.StatusID = *bodyTask.StatusID
	}

	if bodyTask.CreatedAt != nil {
		task.CreatedAt = *bodyTask.CreatedAt
	}

	fmt.Printf("%#v", task)

	_, err = th.TasksStore.UpdateTask(idInt, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task.ToResponse())
}
