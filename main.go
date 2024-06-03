package main

import (
	"context"
	"fmt"
	"internal/database"
	"log"
	"net/http"
	"os"

	"github.com/Elias-Belkheiri/blog_aggregator/controllers"
	"github.com/Elias-Belkheiri/blog_aggregator/models"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()
	ctx := context.Background()

	db, err := sqlx.Connect("postgres", "user=user dbname=db sslmode=disable password=password host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	dbQueries := database.New(db)

	// mux := http.NewServeMux()
	r.Get("/v1/users", models.MiddlewareAuth(controllers.GetUser, ctx, dbQueries))
	r.Post("/v1/users", controllers.AddUserHandler(dbQueries, ctx))

	r.Post("/v1/feeds", models.MiddlewareAuth(controllers.AddFeed, ctx, dbQueries))
	r.Get("/v1/feeds", controllers.GetFeeds(dbQueries, ctx))
	r.Post("/v1/feedFollows", models.MiddlewareAuth(controllers.AddFeedFollows, ctx, dbQueries))
	r.Delete("/v1/feedFollows/{feedFollowsID}", models.MiddlewareAuth(controllers.RemoveFeedFollows, ctx, dbQueries))
	// r.Get("/v1/feed_follows")
	// mux.HandleFunc("POST /v1/users", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.AddUser(w, r, dbQueries, ctx)
	// })
	// mux.HandleFunc("GET /v1/users", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.GetUsers(w, r, dbQueries, ctx)
	// })
	// mux.HandleFunc("GET /v1/user", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.GetUser(w, r, dbQueries, ctx)
	// })
	// mux.HandleFunc("GET /v1/ids/*", controllers.GetId)
	// mux.HandleFunc("GET /v1/readiness", controllers.GetId)
	// mux.HandleFunc("GET /v1/err", errHandler)

	fmt.Println("Listening on port", os.Getenv("PORT"), "...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
