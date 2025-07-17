-- +goose Up
-- +goose StatementBegin
INSERT INTO tasks_categories(id, name)
VALUES 
(1, "category 1"),
(2, "category 2"),
(3, "category 3");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tasks_categories;
-- +goose StatementEnd
