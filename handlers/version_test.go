package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type version struct {
	Version string `json:"version"`
}

func TestGetVersion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()
	golangVersion := os.Getenv("VERSION")

	GetVersion(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	var endpointVersion version

	err = json.Unmarshal(data, &endpointVersion)
	if err != nil {
		log.Fatalf("unexpected error during json Unmarshal: %v", err)
	}

	if endpointVersion.Version != golangVersion {
		t.Errorf("Expected v0.0.1 got %v", endpointVersion.Version)
	}
}
