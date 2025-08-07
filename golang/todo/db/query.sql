-- name: CreateTodo :one
INSERT INTO todos (title) VALUES ($1) RETURNING *;

-- name: ListTodos :many
SELECT * FROM todos ORDER BY created_at DESC;

-- name: UpdateTodo :one
UPDATE todos SET title = $2, completed = $3 WHERE id = $1 RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;
