package controllers

import (
	// "database/sql"
	"context"
	// "database/sql"
	"encoding/json"
	"fmt"
	"internal/database"
	"io"
	"github.com/Elias-Belkheiri/blog_aggregator/utils"
	"net/http"
	"net/url"
	"strconv"
	// "strings"
	"time"
	// "github.com/Elias-Belkheiri/blog_aggregator/models"
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
		utils.RespondWithJSON(w, 200, utils.Test{"All good :)"})
	} else {
		utils.RespondWithJSON(w, 404, utils.Test{"Not found :("})
	}
}

func AddUserHandler(dbQueries *database.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		AddUser(w, r, dbQueries, ctx)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	var user database.CreateUserParams
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Err reading request body")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Err unmarshaling request body")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	// check if the user.name is invalid
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	userCreated, err := dbQueries.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("Err creating user")
		fmt.Println(user.Name)
		fmt.Println(user.ID)
		fmt.Println(user.CreatedAt)
		fmt.Println(user.UpdatedAt)
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	userJson, err := json.Marshal(userCreated)
	if err != nil {
		fmt.Println("Err marshaling user")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(userJson)
}

func GetUsers(w http.ResponseWriter, r *http.Request, dbQueries *database.Queries, ctx context.Context) {
	users, err := dbQueries.GetUsers(ctx)
	if err != nil {
		fmt.Println("Err getting users")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Err marshaling users")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}

	w.Write(resp)
}

func GetUser(w http.ResponseWriter, r *http.Request, user database.User, dbQueries *database.Queries, ctx context.Context) {
	userRetrieved, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Err marshaling user")
		utils.ErrHandler(w, 500, "Internal Server Error")
		return
	}
	w.Write(userRetrieved)
}
