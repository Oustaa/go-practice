package services

import "database/sql"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TasksService struct {
	DB *sql.DB
}

func (th TasksService) GetTasks() ([]Task, error) {
	return []Task{}, nil
}

func (th TasksService) GetTaskById() (Task, error) {
	task := Task{}

	return task, nil
}

func (th TasksService) CreateTask() (Task, error) {
	task := Task{}

	return task, nil
}

func (th TasksService) UpdateTasks() (Task, error) {
	task := Task{}

	return task, nil
}

func (th TasksService) DeleteTasks() error {
	return nil
}
