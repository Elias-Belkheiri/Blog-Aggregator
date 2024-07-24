package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"internal/database"
	"io"
	"net/http"
	"time"

	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Description   string    `xml:"description"`
	Link          string    `xml:"link"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []Item    `xml:"item"`
}

// Item represents an individual item in the RSS feed
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
}

type FeedData struct {
	Feed       database.Feed       	`json:"feed"`
	FeedFollow database.Feedfollow 	`json:"feed_follow"`
}

func AddFeed(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context) {
	var feed database.CreateFeedParams
	var feedFollow database.CreateFeedFollowsParams

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
	// feed.UserID = sql.NullString{String: user.ID, Valid: true}

	feedCreated, err := dbQueries.CreateFeed(ctx, feed)
	if err != nil {
		fmt.Println("Err creating feed")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	feedFollow.ID = uuid.New().String()
	feedFollow.CreatedAt = time.Now()
	feedFollow.UpdatedAt = time.Now()
	feedFollow.UserID = sql.NullString{user.ID, true}
	feedFollow.FeedID = sql.NullString{feed.ID, true}

	feedFollowsCreated, err := dbQueries.CreateFeedFollows(ctx, feedFollow)
	if err != nil {
		fmt.Println("Err creating feedFollows")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	feedJson, err := json.Marshal(FeedData{feedCreated, feedFollowsCreated})
	if err != nil {
		fmt.Println("Err marshaling feed")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(feedJson)
}

func GetFeeds(dbQueries *database.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		feeds, err := dbQueries.GetFeeds(ctx)
	
		if err != nil {
			fmt.Println("Err retrieving feeds")
			utils.ErrHandler(w, 500, "Internal Server Error")
		}
	
		feedsRetrievd, err := json.Marshal(feeds)
		if err != nil {
			fmt.Println("Err marshaling feeds")
			utils.ErrHandler(w, 500, "Internal Server Error")
		}
	
		w.Write(feedsRetrievd)
	}
}

func FetchFeed(url string, feed_id string, dbQueries *database.Queries, ctx context.Context) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Err fetching feed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Err reading the feed body")
		return
	}

	var rss RSSFeed
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Printf("Error parsing the XML: %v\n", err)
		return
	}

	// fmt.Printf("Title: %s\n", rss.Channel.Title)
	// fmt.Printf("Description: %s\n", rss.Channel.Description)
	// fmt.Printf("Link: %s\n", rss.Channel.Link)
	// fmt.Printf("Last Build Date: %s\n", rss.Channel.LastBuildDate)
	// fmt.Println("Items:")
	for _, item := range rss.Channel.Items {
		t, err := time.Parse(item.PubDate, time.RFC3339)
		if err != nil {
			fmt.Println("Err parsing time")
			return
		}
		_, err = dbQueries.CreatePost(ctx, database.CreatePostParams{Title: item.Title, Url: item.Link, Description: sql.NullString{String: "item.Description", Valid: true}, PublishedAt: sql.NullTime{Time: t, Valid: true}, FeedID: feed_id})
		if err != nil {
			fmt.Println("Err creating post")
		}
		// fmt.Printf("  - Title: %s\n", item.Title)
		// fmt.Printf("    Description: %s\n", item.Description)
		// fmt.Printf("    Link: %s\n", item.Link)
		// fmt.Printf("    PubDate: %s\n", item.PubDate)
	}
	fmt.Println("----------------------------------------------------------------------------------------------")
}

func LoopAndFetch(dbQueries *database.Queries, ctx context.Context) {
	var feeds []database.Feed
	var err error

	for {
		feeds, err = dbQueries.GetNextFeedsToFetch(ctx)
		if err != nil {
			fmt.Println("Err getting next feeds to fetch")
			return
		}

		_, err := dbQueries.MarkFeedAsFetched(ctx)
		if err != nil {
			fmt.Println("Err marking feeds as fetched")
			return
		}
	
		for _, feed := range feeds {
			FetchFeed(feed.Url.String, feed.ID, dbQueries, ctx)
		}
	
		time.Sleep(4 * time.Second)
	}
}