package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	handleHealth(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if rec.Header().Get("content-type") != "application/json" {
		t.Fatalf("expected json response, got %q", rec.Header().Get("content-type"))
	}
}

func TestHandleProducts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()

	handleProducts(rec, req)

	var body struct {
		Products []struct {
			ID string `json:"id"`
		} `json:"products"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}

	if len(body.Products) != 4 {
		t.Fatalf("expected 4 demo products, got %d", len(body.Products))
	}
}
