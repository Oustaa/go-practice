package store

import (
	"database/sql"
	"errors"
	"fmt"
)

// type definition for the actuall task
type Task struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description *string      `json:"description,omitempty"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Category    CategoryType `json:"category"`
	Status      StatusType   `json:"status"`
}

type CategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StatusType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TaskResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Category    CategoryType `json:"cateogry"`
	Status      StatusType   `json:"status"`
}

type MySQLTasksService struct {
	DB *sql.DB
}

func NewMySQLTasksService(DB *sql.DB) *MySQLTasksService {
	return &MySQLTasksService{DB}
}

func (th MySQLTasksService) GetTasks() ([]Task, error) {
	var tasks []Task

	query := `
		SELECT t.id, t.title, t.description, t.created_at, t.updated_at, st.name, st.id, cat.name, cat.id FROM tasks as t JOIN tasks_categories as cat
		ON cat.id = t.category_id
		JOIN tasks_statuses as st
		ON st.id = t.status_id
	`

	rows, err := th.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task Task
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt, &task.Status.Name, &task.Status.ID, &task.Category.Name, &task.Category.ID); err != nil {
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

	fmt.Printf("passed id id %d\n", id)

	// should handel the case where description is null, and make join to the two other tables
	query := `
		SELECT t.id, t.title, t.description, t.created_at, t.updated_at, st.name, st.id, cat.name, cat.id FROM tasks as t JOIN tasks_categories as cat
		ON cat.id = t.category_id
		JOIN tasks_statuses as st
		ON st.id = t.status_id
		WHERE t.id = ?;
	`

	err := th.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt, &task.Status.Name, &task.Status.ID, &task.Category.Name, &task.Category.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
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
		fmt.Println(err.Error())
		return task, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(task.Title, task.Status.ID, task.Category.ID)
	if err != nil {
		fmt.Println(err.Error())
		return task, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Failed to get inserted ID:", err)
		return task, err
	}

	createdTask, err := th.GetTaskById(int64(insertedID))
	if err != nil {
		fmt.Printf("%#v", err)
		return Task{}, err
	}

	return *createdTask, nil
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

	_, err := th.DB.Exec(query, task.Title, task.Description, task.Status.ID, task.Category.ID, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
