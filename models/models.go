package models

import (
	"context"
	// "database/sql"
	// "fmt"
	"internal/database"
	"net/http"
	// "strings"
	// "github.com/Elias-Belkheiri/blog_aggregator/utils"
)

type AuthedHandler func(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context)

type ApiConfig struct {
	DB *database.Queries
}

// func MiddlewareAuth(handler AuthedHandler, ctx context.Context, db *database.Queries) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		val, ok := r.Header["Authorization"]
// 		if !ok {
// 			fmt.Println("Missing Auth header")
// 			utils.ErrHandler(w, 400, "Missing Auth header")
// 			return
// 		}
		
// 		authHeader := strings.Split(val[0], " ")
// 		if len(authHeader) != 2 || authHeader[0] != "ApiKey" {
// 			fmt.Println("Invalid Auth header")
// 			utils.ErrHandler(w, 401, "Invalid Auth header")
// 			return
// 		}
	
// 		user, err := db.GetUser(ctx, sql.NullString{authHeader[1], true})
// 		if err != nil {
// 			fmt.Println("Err getting user")
// 			utils.ErrHandler(w, 401, "User not authorized")
// 			return
// 		}

// 		handler(w, r, user, db, ctx)
// 	}
// }
