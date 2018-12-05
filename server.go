package main

import (
	"io"
	"log"
	"net/http"
	"strings"

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
	var metrics []string

	for _, check := range checks {
		metrics = append(metrics, check.getMetric())
	}

	metrics = append(metrics, getMetrics())

	io.WriteString(w, strings.Join(metrics, "\n"))
}
