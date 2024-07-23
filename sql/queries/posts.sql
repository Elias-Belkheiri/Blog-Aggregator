-- name: CreatePost :one
INSERT INTO Posts (
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPostsByUser :many
SELECT * FROM Posts WHERE feed_id IN (
    SELECT feed_id FROM feedFollows WHERE user_id = $1
); 