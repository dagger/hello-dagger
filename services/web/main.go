package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dagger/hello-dagger/internal/catalog"
)

type response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

type productsResponse struct {
	Products []catalog.Product `json:"products"`
	Version  string            `json:"version"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.HandleFunc("GET /products", handleProducts)

	addr := ":" + env("PORT", "8080")
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("demo web service listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, response{
		Message: "hello from the Go demo service",
		Time:    time.Now().UTC().Format(time.RFC3339),
	})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, response{
		Message: "ok",
		Time:    time.Now().UTC().Format(time.RFC3339),
	})
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, productsResponse{
		Products: catalog.FeaturedProducts(),
		Version:  catalog.CatalogVersion,
	})
}

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("content-type", "application/json")
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
