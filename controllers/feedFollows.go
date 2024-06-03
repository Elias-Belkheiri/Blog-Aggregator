package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"io"
	"net/http"
	"strings"
	"time"
	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"github.com/google/uuid"
)

func GetUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context) {
	userFeedFollows, err := dbQueries.GetUserFeedFollows(ctx, sql.NullString{user.ID, true})

	if err != nil {
		fmt.Println("Err getting feedFollows")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	feedFollows, err := json.Marshal(userFeedFollows)
	if err != nil {
		fmt.Println("Err marshaling feedFollows")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(feedFollows)
}

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

func RemoveFeedFollows(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context) {
	pathParts := strings.Split(r.URL.Path, "/")
	feedFollowID := pathParts[3]

	feedFollowDeleted, err := dbQueries.DeleteFeedFollows(ctx, feedFollowID)
	if err != nil {
		fmt.Println("Err deleting feedFollow")
		utils.ErrHandler(w, 500, "Err deleting feedFollow")
		return
	}

	feedFollowJson, err := json.Marshal(feedFollowDeleted)
	if err != nil {
		fmt.Println("Err marshaling feedFollowDeleted")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(feedFollowJson)
}