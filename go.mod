module github.com/Elias-Belkheiri/blog_aggregator

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	internal/database v0.0.0
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/go-chi/chi/v5 v5.0.12 // indirect
	github.com/go-chi/render v1.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.25.10 // indirect
)

replace internal/database => ./internal/database

// replace controllers => ./controllers
