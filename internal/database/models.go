// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Apikey    sql.NullString
}
