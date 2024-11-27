package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	golangVersion := os.Getenv("VERSION")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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
