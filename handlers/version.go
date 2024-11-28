package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	golangVersion := os.Getenv("VERSION")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["version"] = golangVersion
	JsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	_, err = w.Write(JsonResp)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
