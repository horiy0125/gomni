package handler

import (
	"encoding/json"
	"net/http"
)

type ApiHealth struct {
	Status  string
	Message string
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	apiHealth := ApiHealth{
		Status:  "OK",
		Message: "Working normally",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiHealth)
}
