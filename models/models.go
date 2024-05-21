package models

import "internal/database"

type Test struct {
	Status string `json:"status"`
}

type Err struct {
	Err string `json:"error"`
}

type ApiConfig struct {
	DB *database.Queries
}