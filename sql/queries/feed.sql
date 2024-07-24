-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST, last_fetched_at ASC LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING *;