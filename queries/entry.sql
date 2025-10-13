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
