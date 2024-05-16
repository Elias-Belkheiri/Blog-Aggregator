package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"io"
	"context"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
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
	mux.HandleFunc("POST /v1/users", func (w http.ResponseWriter, r *http.Request) {
		addUser(w, r, dbQueries, ctx) })
	mux.HandleFunc("GET /v1/ids/*", getId)
	mux.HandleFunc("GET /v1/readiness", readAble)
	// mux.HandleFunc("GET /v1/err", errHandler)
	
	fmt.Println("Listening on port", os.Getenv("PORT"), "...")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), mux))
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