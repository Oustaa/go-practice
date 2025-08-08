-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks
ADD COLUMN user_id INT,
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks
DROP FOREIGN KEY fk_user,
DROP COLUMN user_id;
-- +goose StatementEnd

