package api

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

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

	type TaskToUpdateInterfce struct {
		Title       *string
		Description *string
		CreatedAt   *string
		UpdatedAt   *string
		CategoryID  *int
		StatusID    *int
	}

	bodyTask := &TaskToUpdateInterfce{}
	json.NewDecoder(r.Body).Decode(&bodyTask)
	defer r.Body.Close()

	if bodyTask.Title != nil {
		task.Title = *bodyTask.Title
	}

	if bodyTask.Description != nil {
		task.Title = *bodyTask.Description
	}
}
