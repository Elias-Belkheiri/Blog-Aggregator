-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at NULL FIRST, last_fetched_at ASC;