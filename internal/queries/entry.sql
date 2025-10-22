-- name: FindAllEntries :many
SELECT *
FROM entry;

-- name: FindEntryById :one
SELECT * 
FROM entry 
WHERE id = ?
LIMIT 1;

-- name: FinishEntry :exec
UPDATE entry
SET finished=1
WHERE id = ?;

-- name: DeleteEntry :exec
DELETE FROM entry 
WHERE id = ?;

-- name: AddEntry :one
INSERT INTO entry (
    task_id
) VALUES (
    ?
) returning *;

-- name: ClearEntries :exec
DELETE FROM entry;

-- name: CalculateTaskTime :one
SELECT SUM(end_timestamp - start_timestamp) 
FROM entry 
WHERE task_id = ?;
