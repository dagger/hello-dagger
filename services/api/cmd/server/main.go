package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dagger/demo-monorepo/services/api/internal/catalog"
)

type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.HandleFunc("GET /api/products", handleProducts)

	addr := ":" + env("PORT", "8080")
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("api listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Status: "ok",
		Time:   time.Now().UTC().Format(time.RFC3339),
	})
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"products": catalog.FeaturedProducts(),
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("write response: %v", err)
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
