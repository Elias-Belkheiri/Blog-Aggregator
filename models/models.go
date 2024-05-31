package models

import (
	"context"
	"database/sql"
	"fmt"
	"internal/database"
	"net/http"
	"strings"
	"github.com/Elias-Belkheiri/blog_aggregator/utils"
)

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)

type ApiConfig struct {
	DB *database.Queries
}

func MiddlewareAuth(handler AuthedHandler, ctx context.Context, db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Header["Authorization"]
		if !ok {
			fmt.Println("Missing Auth header")
			utils.ErrHandler(w, r, http.StatusUnauthorized, "Missing Auth header")
			return
		}
		
		authHeader := strings.Split(val[0], " ")
		if len(authHeader) != 2 || authHeader[0] != "ApiKey" {
			fmt.Println("Invalid Auth header")
			utils.ErrHandler(w, r, http.StatusUnauthorized, "Invalid Auth header")
			return
		}
	
		user, err := db.GetUser(ctx, sql.NullString{authHeader[1], true})
		if err != nil {
			fmt.Println("Err getting user")
			utils.ErrHandler(w, r, http.StatusNotFound, "User not authorized")
			return
		}

		handler(w, r, user)
	}
}
