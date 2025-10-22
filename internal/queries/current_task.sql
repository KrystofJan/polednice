-- name: GetCurrentTask :one
SELECT *
FROM current_task
LIMIT 1;
