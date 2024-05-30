-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, apikey)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE apikey = $1;

-- name: GetUsers :many
SELECT * FROM users;