package api

import (
	"encoding/json"
	"net/http"

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

	json.NewEncoder(w).Encode(tasks)
}

func (th Taskshandler) PostTasks(w http.ResponseWriter, r *http.Request) {
	// Get The Data from the body
	var task store.Task // it is not the ideal way to get the Task struct
	json.NewDecoder(r.Body).Decode(&task)
	defer r.Body.Close()

	// Create the task
	th.TasksStore.CreateTask(task)

	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(task)
}
