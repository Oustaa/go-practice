-- +goose Up
-- +goose StatementBegin
INSERT INTO tasks_categories(name)
values
("category 1"),
("category 2");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tasks_categories
WHERE name in ("category 1", "category 2");
-- +goose StatementEnd
