package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"io"
	"net/http"
	"time"
	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func AddFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	var feed database.CreateFeedParams
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Err reading request body")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	err = json.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Err unmarshalling request body")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	if feed.Name == "" {
		fmt.Println("missing name attribute")
		utils.ErrHandler(w, 400, "Invalid name attribute")
		return
	}
	feed.ID = uuid.New().String()
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = time.Now()
	feed.UserID = sql.NullString{String: user.ID, Valid: true}

	feedCreated, err := json.Marshal(feed)
	if err != nil {
		fmt.Println("Err marshaling feed")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(feedCreated)
	// utils.RespondWithJSON(w, 200, )
	// w.Write(feedRetrieved)
}