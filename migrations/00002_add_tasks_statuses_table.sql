-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks_statuses (
    id INT AUTO_INCREMENT PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks_statuses;
-- +goose StatementEnd