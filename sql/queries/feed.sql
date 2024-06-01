-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;