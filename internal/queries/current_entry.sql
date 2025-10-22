-- name: GetCurrentEntry :one
SELECT *
FROM current_entry
LIMIT 1;
