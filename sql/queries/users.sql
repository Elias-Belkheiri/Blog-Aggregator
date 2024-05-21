-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, apikey,  name)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE apikey = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;