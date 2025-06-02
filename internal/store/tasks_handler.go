package store

import (
	"database/sql"
)

// type definition for the actuall task
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CategoryID  int    `json:"category_id"`
	StatusID    int    `json:"status_id"`
}

type MySQLTasksService struct {
	DB *sql.DB
}

func NewMySQLTasksService(DB *sql.DB) *MySQLTasksService {
	return &MySQLTasksService{DB}
}

func (th MySQLTasksService) GetTasks() ([]Task, error) {
	var tasks []Task

	query := "SELECT id, title, created_at, updated_at FROM tasks"

	rows, err := th.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task Task
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (th MySQLTasksService) GetTaskById() (Task, error) {
	task := Task{}

	return task, nil
}

func (th MySQLTasksService) CreateTask(task Task) (Task, error) {
	stmt, err := th.DB.Prepare(`
		INSERT INTO tasks(title, status_id, category_id) VALUES(?, ?, ?);`)
	if err != nil {
		return task, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.StatusID, task.CategoryID)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (th MySQLTasksService) UpdateTasks() (Task, error) {
	task := Task{}

	return task, nil
}

func (th MySQLTasksService) DeleteTasks() error {
	return nil
}
