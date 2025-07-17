-- +goose Up
-- +goose StatementBegin
INSERT INTO tasks_statuses(id, name, category_id)
VALUES
(1, "sts 1 1", 1),
(2, "sts 2 1", 1),
(3, "sts 1 2", 2),
(4, "sts 2 2", 2),
(5, "sts 1 3", 3),
(6, "sts 2 3", 3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tasks_statuses;
-- +goose StatementEnd
