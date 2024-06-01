package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"io"
	"net/http"
	"time"

	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"github.com/google/uuid"
)


func AddFeedFollows(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context) {
	var extractedFeedFollow database.CreateFeedFollowsParams
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Err reading request body")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	err = json.Unmarshal(body, &extractedFeedFollow)
	if err != nil {
		fmt.Println("Err unmarshalling request body")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	if !extractedFeedFollow.FeedID.Valid {
		fmt.Println("Missing feedID")
		utils.ErrHandler(w, 400, "Missing feedID")
		return
	}

	extractedFeedFollow.ID = uuid.New().String()
	extractedFeedFollow.CreatedAt = time.Now()
	extractedFeedFollow.UpdatedAt = time.Now()
	extractedFeedFollow.UserID = sql.NullString{user.ID, true}

	feedFollowsCreated, err := dbQueries.CreateFeedFollows(ctx, extractedFeedFollow)
	if err != nil {
		fmt.Println("Err creating feedFollows")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	feedFollows, err := json.Marshal(feedFollowsCreated)
	if err != nil {
		fmt.Println("Err marshaling feedFollows")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(feedFollows)
}