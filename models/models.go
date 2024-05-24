package models

import (
	"internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

type Test struct {
	Status string `json:"status"`
}

type Err struct {
	Err string `json:"error"`
}

type ApiConfig struct {
	DB *database.Queries
}

func (cfg *ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    ///
}
