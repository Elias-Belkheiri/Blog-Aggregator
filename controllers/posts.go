package controllers

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	// "database/sql"
// 	"internal/database"
// 	"net/http"

// 	"github.com/Elias-Belkheiri/blog_aggregator/utils"
// )

// // "github.com/Elias-Belkheiri/blog_aggregator/utils"
// // "internal/database"

// // Create a func to addPosts

// type Post struct {
// 	Title 		string `json:"title"`
// 	Url  		string `json:"url"`
// 	Description string `json:"description"`
// 	PublishedAt string `json:"published_at"`
// }

// func CreatePost(dbQueries *database.Queries, ctx context.Context, item Item, feed database.Feed) {
// 	t, err := time.Parse(time.RFC1123Z, item.PubDate)
// 	if err != nil {
// 		fmt.Println("Err parsing time")
// 		return
// 	}
	
// 	_, err = dbQueries.CreatePost(ctx, database.CreatePostParams{
// 		Title:       		item.Title,
// 		Description: 		sql.NullString{item.Description, true},
// 		Url:        		item.Link,
// 		PublishedAt:     	sql.NullTime{t, true},
// 		FeedID:      		feed.ID,
// 	})

// 	if err != nil {
// 		fmt.Println("Err creating post")
// 		return
// 	}
// }

// func GetPostsByUser (w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context) {
// 	postsRetrieved, err := dbQueries.GetPostsByUser(ctx, sql.NullString{user.ID, true})
// 	var posts []Post
// 	if err != nil {
// 		utils.ErrHandler(w, 500, "Err getting posts")
// 		return
// 	}

// 	for _, post := range postsRetrieved {
// 		posts = append(posts, Post{
// 			Title: post.Title,
// 			Url: post.Url,
// 			Description: post.Description.String,
// 			PublishedAt: post.PublishedAt.Time.String(),
// 		})
// 	}

// 	utils.RespondWithJSON(w, 200, posts)
// }