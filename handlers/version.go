package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	golangVersion := os.Getenv("VERSION")
	fmt.Fprint(w, golangVersion)
}
