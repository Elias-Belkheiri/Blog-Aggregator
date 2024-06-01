// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed-follows.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createFeedFollows = `-- name: CreateFeedFollows :one

INSERT INTO feedFollows (id, user_id, feed_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, feed_id, created_at, updated_at
`

type CreateFeedFollowsParams struct {
	ID        string
	UserID    sql.NullString
	FeedID    sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFeedFollows(ctx context.Context, arg CreateFeedFollowsParams) (Feedfollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollows,
		arg.ID,
		arg.UserID,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Feedfollow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedFollows = `-- name: GetFeedFollows :many
SELECT id, user_id, feed_id, created_at, updated_at FROM feedFollows
`

func (q *Queries) GetFeedFollows(ctx context.Context) ([]Feedfollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feedfollow
	for rows.Next() {
		var i Feedfollow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
