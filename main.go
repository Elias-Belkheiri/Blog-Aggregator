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
	r := chi.NewRouter()
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") 
	}

	connString := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"), "localhost", "5432")

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}

	dbQueries := database.New(db)
	
	r.Get("/v1/users", models.MiddlewareAuth(controllers.GetUser, ctx, dbQueries))
	r.Post("/v1/register", controllers.AddUserHandler(dbQueries, ctx))
	r.Post("/v1/login", controllers.LogUserInHandler(dbQueries, ctx))
	
	r.Post("/v1/feeds", models.MiddlewareAuth(controllers.AddFeed, ctx, dbQueries))
	// r.Get("/v1/feeds", controllers.GetFeeds(dbQueries, ctx))

	r.Get("/v1/posts", models.MiddlewareAuth(controllers.GetPostsByUser, ctx, dbQueries))

	go controllers.LoopAndFetch(dbQueries, ctx)

	fmt.Println("Listening on port", os.Getenv("PORT"), "...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
