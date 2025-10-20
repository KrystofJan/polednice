-- name: SwitchCurrentTask :one
INSERT INTO current_task (
    task_id
) VALUES (
    (SELECT id FROM current_task LIMIT 1)
) returning *;
