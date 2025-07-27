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

type CreateTaskBody struct {
	Title       string  `json:"title"`
	Status_id   int     `json:"status_id"`
	Category_id int     `json:"category_id"`
	Description *string `json:"description"`
}

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
	json.NewEncoder(w).Encode(task)
}

func (th Taskshandler) PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Get The Data from the body
	var taskBody CreateTaskBody
	json.NewDecoder(r.Body).Decode(&taskBody)
	defer r.Body.Close()

	task := store.Task{
		Title: taskBody.Title,
		Category: store.CategoryType{
			ID: taskBody.Category_id,
		},
		Status: store.StatusType{
			ID: taskBody.Status_id,
		},
		Description: taskBody.Description,
	}

	fmt.Println()
	fmt.Printf("task to create body: %#v", task)
	fmt.Println()

	// Create the task
	createdTask, err := th.TasksStore.CreateTask(task)

	if err != nil {
		fmt.Printf("Error Creating the task: %#v", err)
	}

	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(createdTask)
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
		// task.Description = sql.NullString{String: *bodyTask.Description, Valid: true}
		task.Description = bodyTask.Description
	}

	if bodyTask.CategoryID != nil {
		task.Category.ID = *bodyTask.CategoryID
	}

	if bodyTask.StatusID != nil {
		task.Status.ID = *bodyTask.StatusID
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
	json.NewEncoder(w).Encode(task)
}
