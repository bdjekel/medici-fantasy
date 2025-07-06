package main

import (
	"net/http"
	"time"
)

// healthHandler handles health check requests
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	response := Response{
		Status:    "healthy",
		Timestamp: time.Now(),
	}
	
	writeJSON(w, http.StatusOK, response)
}