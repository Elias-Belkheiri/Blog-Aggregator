package utils

import (
	"encoding/json"
	"net/http"
	"github.com/Elias-Belkheiri/blog_aggregator/models"
)

func ReadAble(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, models.Test{Status: "ok"})
}

func ErrHandler(w http.ResponseWriter, r *http.Request, err int, description string) {
	RespondWithJSON(w, err, models.Err{description})
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, "err marshaling", http.StatusInternalServerError)
	}

	w.Write([]byte(response))
}