-- name: FindAllTasks :many
SELECT * 
FROM task;

-- name: FindTaskById :one
SELECT * 
FROM task 
WHERE id = ?
LIMIT 1;

-- name: FindTaskByName :one
SELECT * 
FROM task 
WHERE name = ?
LIMIT 1;

-- name: FinishTask :exec
UPDATE task
SET finished = 1
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM task 
WHERE id = ?;

-- name: AddTask :one
INSERT INTO task (
    name
) VALUES (
    ? 
) returning *;
