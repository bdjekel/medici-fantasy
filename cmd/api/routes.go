package main

import (
	"net/http"
	"strings"
	"time"
)

// handleRoutes handles all routing logic
func handleRoutes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	
	switch {
	case path == "/health":
		healthHandler(w, r)
	case strings.HasPrefix(path, "/api/v1"):
		apiHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

// apiHandler handles API v1 routes
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement API v1 routing logic
	response := Response{
		Message:   "API v1 endpoint - not yet implemented",
		Status:    "success",
		Timestamp: time.Now(),
		Version:   "v1",
	}
	
	writeJSON(w, http.StatusOK, response)
}

