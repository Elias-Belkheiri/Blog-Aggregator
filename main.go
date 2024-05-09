package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Test struct {
	Status string `json:"status"`
}

type Err struct {
	Err string `json:"error"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/ids/*", getId)
	mux.HandleFunc("GET /v1/readiness", readAble)
	mux.HandleFunc("GET /v1/err", errHandler)
	
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

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 500, Err{"Internal server error"})
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