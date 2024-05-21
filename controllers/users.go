package controllers

import (
	// "database/sql"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"io"
	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"github.com/Elias-Belkheiri/blog_aggregator/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func GetId(w http.ResponseWriter, r *http.Request) {
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
		utils.RespondWithJSON(w, 200, models.Test{"All good :)"})
	} else {
		utils.RespondWithJSON(w, 404, models.Test{"Not found :("})
	}
}

func AddUser(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	var user database.CreateUserParams
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Err reading request body")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Err unmarshaling request body")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	userCreated, err := dbQueries.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("Err creating user")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	userJson, err := json.Marshal(userCreated)
	if err != nil {
		fmt.Println("Err marshaling user")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	w.Write(userJson)
}

func GetUsers(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	users, err := dbQueries.GetUsers(ctx)
	if err != nil {
		fmt.Println("Err getting users")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Err marshaling users")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Write(resp)
}

func GetUser(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	// var user database.User

	authHeader := strings.Split(r.Header["Authorization"][0], " ")
	if len(authHeader) != 2 || authHeader[0] != "ApiKey" {
		fmt.Println("Invalid Auth header")
		utils.ErrHandler(w, r, http.StatusUnauthorized, "Invalid Auth header")
		return
	}

	user, err := dbQueries.GetUser(ctx, sql.NullString{authHeader[1], true})
	if err != nil {
		fmt.Println("Err getting user")
		utils.ErrHandler(w, r, http.StatusNotFound, "User not found")
		return
	}

	userRetrieved, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Err marshaling user")
		utils.ErrHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	fmt.Println(authHeader[1])
	fmt.Println(userRetrieved)

	w.Write(userRetrieved)
}
