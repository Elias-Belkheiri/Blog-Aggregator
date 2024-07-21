// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID            string
	Name          string
	Url           sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
	LastFetchedAt sql.NullTime
}

type Feedfollow struct {
	ID        string
	UserID    sql.NullString
	FeedID    sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Apikey    sql.NullString
}
