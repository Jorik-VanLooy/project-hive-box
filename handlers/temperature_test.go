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

type response struct {
	Temperature float32 `json:"temperature"`
}

func TestGetTemperature(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/temperature", nil)
	w := httptest.NewRecorder()

	GetTemperature(w, req)
	res := w.Result()
	defer res.Body.Close()

	var data openweather
	apiKey := os.Getenv("OpenWeatherApiKey")

	testData, _ := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, apiKey))
	bodyBytes, _ := io.ReadAll(testData.Body)
	err := json.Unmarshal([]byte(string(bodyBytes)), &data)
	if err != nil {
		log.Fatal(err)
	}

	var responseEndpoint response

	strTemp := fmt.Sprintf("%.2f", data.MainStats.Temp-kelvin)

	tempEndPoint, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = json.Unmarshal([]byte(string(tempEndPoint)), &responseEndpoint)
	if err != nil {
		t.Fatalf("error in the json unmarshal from endpoint: %v", err)
	}
	strEndPoint := fmt.Sprintf("%.2f", responseEndpoint.Temperature)

	if strTemp != strEndPoint {
		t.Errorf("Expected %s got %s", strTemp, strEndPoint)
	}
}
