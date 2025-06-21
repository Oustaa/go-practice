-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    status_id INT NOT NULL, -- DONE, IN PROGRESS, TODO ( Defined later by users )
    category_id INT NOT NULL, -- STUDY, WORK, PERSONAL ( Defined later by users )
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (status_id) REFERENCES tasks_statuses (id),
    FOREIGN KEY (category_id) REFERENCES tasks_categories (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
