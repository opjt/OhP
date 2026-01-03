-- name: CreateNotification :one
INSERT INTO notifications (
    endpoint_id,
    user_id,
    body
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateStatusNotification :exec
UPDATE notifications
SET status = $2
WHERE id = $1;

-- name: FindNotificationByUserID :many
SELECT 
    n.*,
    e.name as endpoint_name
FROM notifications n
JOIN endpoints e ON n.endpoint_id = e.id
WHERE n.user_id = $1;
