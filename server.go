package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/metrics", MetricsHandler)
	http.Handle("/", r)

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(checks)
}
