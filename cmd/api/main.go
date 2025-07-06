package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Response represents a standard API response
type Response struct {
	Message   string    `json:"message,omitempty"`
	Status    string    `json:"status,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Version   string    `json:"version,omitempty"`
}

// TODO: Add apiConfig. Below structure is copy pasted from chirpy (boot.dev project)
// type apiConfig struct {
//  fileserverHits atomic.Int32
// 	Db database.Queries
// 	platform string
// 	secret string
// 	polkaKey string
// }


// Server represents the HTTP server
type Server struct {
	server *http.Server
}

// NewServer creates a new server instance
func NewServer(port string) *Server {
	server := &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      http.HandlerFunc(handleRoutes),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
	
	return server
}

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

// apiHandler handles API requests
func apiHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	
	switch path {
	case "/api/v1":
		apiRootHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

// apiRootHandler handles API root requests
func apiRootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	response := Response{
		Message:   "API is running",
		Version:   "1.0.0",
		Timestamp: time.Now(),
	}
	
	writeJSON(w, http.StatusOK, response)
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	log.Printf("Starting server on port %s", s.server.Addr)
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func main() {
	const filepathRoot = "."
	const env = "dev"
	const logLevel = "debug"
	const readTimeout = 15 * time.Second
	const writeTimeout = 15 * time.Second
	const idleTimeout = 60 * time.Second
	
	godotenv.Load()

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Create new server
	server := NewServer(port)
	
	// Start server in a goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// Shutdown server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	
	log.Println("Server exited")
}
