-- name: CreateEndpoint :one
INSERT INTO endpoints (
    user_id,
    name,
    token
) VALUES (
    $1, 
    $2, 
    $3
)
RETURNING *;

-- name: FindEndpointByUserID :many
SELECT * FROM endpoints
WHERE user_id = $1;

-- name: FindEndpointByToken :one
SELECT * FROM endpoints
WHERE token = $1;

-- name: DeleteEndpointByToken :exec
DELETE FROM endpoints
WHERE token = $1
  AND user_id = $2;

-- name: UpdateEndpointMute :exec
UPDATE endpoints
SET notification_enabled = false, 
  notification_disabled_at = $2
WHERE token = $1;

-- name: UpdateEndpointUnmute :exec
UPDATE endpoints
SET notification_enabled = true, 
  notification_disabled_at = null
WHERE token = $1;