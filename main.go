package main

import (
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	fmt.Printf("Listening on port %s ...", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), mux))
}