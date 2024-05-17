package main

import (
	// "database/sql"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// "github.com/lib/pq"
)

type Test struct {
	Status string `json:"status"`
}

type Err struct {
	Err string `json:"error"`
}

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// dbURL := os.Getenv("POSTGRES_URI")
	ctx := context.Background()
	// fmt.Println(dbURL)

	// connStr := "user=myuser dbname=mydb password=password sslmode=disable"
	db, err := sqlx.Connect("postgres", "user=hmeda dbname=blogy sslmode=disable password=password host=localhost")
	if err != nil {
		log.Fatalln(err)
	}
	dbQueries := database.New(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/users", func(w http.ResponseWriter, r *http.Request) {
		addUser(w, r, dbQueries, ctx)
	})
	mux.HandleFunc("GET /v1/users", func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, dbQueries, ctx)
	})
	mux.HandleFunc("GET /v1/ids/*", getId)
	mux.HandleFunc("GET /v1/readiness", readAble)
	// mux.HandleFunc("GET /v1/err", errHandler)

	fmt.Println("Listening on port", os.Getenv("PORT"), "...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, "err marshaling", http.StatusInternalServerError)
	}

	w.Write([]byte(response))
}

func readAble(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, Test{Status: "ok"})
}

func errHandler(w http.ResponseWriter, r *http.Request, err int, description string) {
	respondWithJSON(w, err, Err{description})
}

func getId(w http.ResponseWriter, r *http.Request) {
	uri, err := url.Parse(r.RequestURI)

	if err != nil {
		fmt.Println("Err parsing uri")
	}

	id, err := strconv.Atoi(uri.Path[len("/ids/"):])
	if err != nil {
		fmt.Println("err casting id to int")
	}
	fmt.Printf("id: -%d-", id)
	if id > 10 {
		respondWithJSON(w, 200, Test{"All good :)"})
	} else {
		respondWithJSON(w, 404, Test{"Not found :("})
	}
}

func addUser(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	var user database.CreateUserParams
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Err reading request body")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Err unmarshaling request body")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	userCreated, err := dbQueries.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("Err creating user")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	userJson, err := json.Marshal(userCreated)
	if err != nil {
		fmt.Println("Err marshaling user")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	w.Write(userJson)
}

func getUsers(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	users, err := dbQueries.GetUsers(ctx)
	if err != nil {
		fmt.Println("Err getting users")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Err marshaling users")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Write(resp)
}

func getUser(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	// var user database.User

	authHeader := strings.Split(r.Header["Authorization"][0], " ")
	if len(authHeader) != 2 || authHeader[0] != "ApiKey" {
		fmt.Println("Invalid Auth header")
		errHandler(w, r, http.StatusUnauthorized, "Invalid Auth header")
		return
	}

	user, err := dbQueries.GetUser(ctx, sql.NullString{authHeader[1], true})
	if err != nil {
		fmt.Println("Err getting user")
		errHandler(w, r, http.StatusNotFound, "User not found")
		return
	}

	userRetrieved, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Err marshaling user")
		errHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	fmt.Println(authHeader[1])
	fmt.Println(userRetrieved)

	w.Write(userRetrieved)
}
