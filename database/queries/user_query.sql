-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
	name, nis, password, created_at, updated_at
) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
