-- name: UpsertUserByEmail :one
INSERT INTO users (email, updated_at)
VALUES ($1, now())
ON CONFLICT (email) 
DO UPDATE SET 
    updated_at = EXCLUDED.updated_at
RETURNING id, email, created_at, updated_at;