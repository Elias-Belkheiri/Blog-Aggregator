-- name: CreateFeedFollows :one

INSERT INTO feedFollows (id, user_id, feed_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feedFollows;