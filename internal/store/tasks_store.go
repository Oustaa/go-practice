package store

import (
	"database/sql"
	"errors"
)

// type definition for the actuall task
type Task struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"-"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	CategoryID  int            `json:"category_id"`
	StatusID    int            `json:"status_id"`
}

type TaskResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CategoryID  int    `json:"category_id"`
	StatusID    int    `json:"status_id"`
}

func (t *Task) ToResponse() *TaskResponse {
	return &TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description.String,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		CategoryID:  t.CategoryID,
		StatusID:    t.StatusID,
	}
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

func (th MySQLTasksService) GetTaskById(id int64) (*Task, error) {
	task := &Task{}

	// should handel the case where description is null, and make join to the two other tables
	query := "SELECT id, title, description, category_id, status_id, created_at, updated_at FROM tasks where id = ?"

	err := th.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.CategoryID, &task.StatusID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // or return a custom 404 error
		}
		return nil, err
	}

	return task, nil
}

func (th MySQLTasksService) CreateTask(task Task) (Task, error) {
	// this is in a must need for validations
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

func (th MySQLTasksService) DeleteTasks(id int64) error {
	query := "DELETE FROM tasks WHERE id = ?"

	_, err := th.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (th MySQLTasksService) UpdateTask(id int64, task *Task) (*Task, error) {
	query := "UPDATE tasks set title = ?, description = ?, status_id = ?, category_id = ? WHERE id = ?"

	_, err := th.DB.Exec(query, task.Title, task.Description, task.StatusID, task.CategoryID, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
