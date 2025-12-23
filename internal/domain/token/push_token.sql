-- name: UpsertToken :one
INSERT INTO push_tokens (
    user_id,
    p256dh_key,
    auth_key,
    endpoint
)
VALUES ($1, $2, $3, $4)
ON CONFLICT (user_id, p256dh_key, auth_key, endpoint)
DO UPDATE SET
    is_active = true
RETURNING user_id;

-- name: DeleteToken :exec
DELETE FROM push_tokens
WHERE endpoint = $1 
    AND p256dh_key = $2 
    AND auth_key = $3;