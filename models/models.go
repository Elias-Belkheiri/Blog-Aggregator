package models

import (
	"context"
	"fmt"
	"internal/database"
	"net/http"
	"strings"

	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"github.com/golang-jwt/jwt"
)

type AuthedHandler func(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context)

type ApiConfig struct {
	DB *database.Queries
}

func MiddlewareAuth(handler AuthedHandler, ctx context.Context, db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Header["Authorization"]
		if !ok {
			fmt.Println("Missing Auth header")
			utils.ErrHandler(w, 400, "Missing Auth header")
			return
		}
		
		authHeader := strings.Split(val[0], " ")
		if len(authHeader) != 2 || authHeader[0] != "Bearer" {
			fmt.Println("Invalid Auth header")
			utils.ErrHandler(w, 401, "Invalid Auth header")
			return
		}
	
		token, err := utils.VerifyToken(authHeader[1])
		if err != nil {
			fmt.Println("Err verifying token")
			utils.ErrHandler(w, 401, "Invalid token")
			return
		}
		user, err := db.GetUser(ctx, token.Claims.(jwt.MapClaims)["username"].(string))
		if err != nil {
			fmt.Println("Err getting user")
			utils.ErrHandler(w, 401, "User not authorized")
			return
		}

		handler(w, r, user, db, ctx)
	}
}
