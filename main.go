package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horiy0125/gomni/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/health_check", handler.HealthCheckHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
