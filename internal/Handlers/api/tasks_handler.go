package api

import (
	"encoding/json"
	"log"
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

func (th Taskshandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.TasksStore.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (th Taskshandler) PostTask(w http.ResponseWriter, r *http.Request) {
	var task store.Task
	json.NewDecoder(r.Body).Decode(&task)
	defer r.Body.Close()

	createTask, err := th.TasksStore.CreateTask(&task)
	if err != nil {
		log.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(createTask)
}

func (th Taskshandler) PutTask(w http.ResponseWriter, r *http.Request) {
	paramsTaskID := chi.URLParam(r, "id")

	if paramsTaskID == "" {
		http.NotFound(w, r)
		return
	}

	workoutId, err := strconv.ParseInt(paramsTaskID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var task store.TaskUpdate
	json.NewDecoder(r.Body).Decode(&task)

	updatedTask, err := th.TasksStore.UpdateTask(workoutId, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	json.NewEncoder(w).Encode(updatedTask)
}
