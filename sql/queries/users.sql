-- name: CreateUser :one
INSERT INTO users (id, username, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1;

-- name: GetUsers :many
SELECT * FROM users;
