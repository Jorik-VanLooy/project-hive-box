package main

import (
	"net/http"

	"github.com/Jorik-VanLooy/project-hive-box/handlers"
)

func main() {
	http.HandleFunc("/version", handlers.GetVersion)
	http.HandleFunc("/temperature", handlers.GetTemperature)
	http.HandleFunc("/health", handlers.GetHealth)
	http.ListenAndServe(":3333", nil)
}
