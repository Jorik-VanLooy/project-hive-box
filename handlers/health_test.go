package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type health struct {
	Healthy string `json:"message"`
}

func TestGetHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	GetHealth(w, req)
	res := w.Result()
	defer res.Body.Close()

	var Health health

	health, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = json.Unmarshal(health, &Health)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Health.Healthy != "healthy" {
		t.Errorf("expected healthy got %v", Health.Healthy)
	}
}
