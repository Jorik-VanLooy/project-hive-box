package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetTemperature(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/temperature", nil)
	w := httptest.NewRecorder()

	GetTemperature(w, req)
	res := w.Result()
	defer res.Body.Close()

	var data openweather
	apiKey := os.Getenv("OpenWeatherApiKey")

	response, _ := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, apiKey))
	bodyBytes, _ := io.ReadAll(response.Body)
	err := json.Unmarshal([]byte(string(bodyBytes)), &data)
	if err != nil {
		log.Fatal(err)
	}

	strTemp := fmt.Sprintf("%.2f", data.MainStats.Temp-kelvin)

	tempEndPoint, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	if strTemp != string(tempEndPoint) {
		t.Errorf("Expected v0.0.1 got %v", strTemp)
	}
}
