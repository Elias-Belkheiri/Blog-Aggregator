module github.com/Elias-Belkheiri/blog_aggregator

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	internal/database v0.0.0
)

replace internal/database => ./internal/database

// replace controllers => ./controllers
