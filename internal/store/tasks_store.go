package store

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type NullStringJSON struct {
	sql.NullString
}

func (nsj NullStringJSON) MarshalJSON() ([]byte, error) {
	if nsj.Valid {
		return json.Marshal(nsj.String)
	}
	return json.Marshal("")
}

func (nsj *NullStringJSON) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		nsj.String = *s
		nsj.Valid = true
	} else {
		nsj.Valid = false
	}
	return nil
}

type Task struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description NullStringJSON `json:"description"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	CategoryID  int            `json:"category_id"`
	StatusID    int            `json:"status_id"`
}

type TaskUpdate struct {
	Title       *string         `json:"title"`
	Description *NullStringJSON `json:"description"`
	CategoryID  *int            `json:"category_id"`
	StatusID    *int            `json:"status_id"`
}

type MySQLTasksService struct {
	DB *sql.DB
}

func NewMySQLTasksService(DB *sql.DB) *MySQLTasksService {
	return &MySQLTasksService{DB}
}

func (th MySQLTasksService) GetTasks() ([]Task, error) {
	var tasks []Task

	query := "SELECT id, title, created_at, updated_at, description FROM tasks"

	rows, err := th.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task Task
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.CreatedAt, &task.UpdatedAt, &task.Description); err != nil {
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

func (th MySQLTasksService) CreateTask(task *Task) (*Task, error) {
	insertQuery := `
		INSERT INTO tasks(title, status_id, category_id, description)
		VALUES (?, ?, ?, ?);`

	result, err := th.DB.Exec(insertQuery, task.Title, task.StatusID, task.CategoryID, task.Description.NullString)
	if err != nil {
		return nil, fmt.Errorf("error while creating the task: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	selectQuery := `
		SELECT id, title, status_id, category_id, created_at, description
		FROM tasks
		WHERE id = ?;`

	row := th.DB.QueryRow(selectQuery, id)

	var createdTask Task
	err = row.Scan(&createdTask.ID, &createdTask.Title, &createdTask.StatusID, &createdTask.CategoryID, &createdTask.CreatedAt, &createdTask.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the inserted task: %v", err)
	}

	return &createdTask, nil
}

func (th MySQLTasksService) UpdateTask(id int64, task TaskUpdate) (TaskUpdate, error) {

	return task, nil
}

func (th MySQLTasksService) DeleteTasks() error {
	return nil
}
