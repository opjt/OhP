-- name: CreateNotification :one
INSERT INTO notifications (
    endpoint_id,
    body
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateStatusNotification :exec
UPDATE notifications
SET status = $2
WHERE id = $1;