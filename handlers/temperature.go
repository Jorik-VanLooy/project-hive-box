package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	lat    = 51.166827
	lon    = 4.702270
	kelvin = 272.15
)

type coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type mainStats struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type wind struct {
	Speed float32 `json:"speed"`
	Deg   int     `json:"deg"`
}

type clouds struct {
	All int `json:"all"`
}

type sys struct {
	Types   int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type openweather struct {
	Coord      coord     `json:"coord"`
	Weather    []weather `json:"weather"`
	Base       string    `json:"base"`
	MainStats  mainStats `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       wind      `json:"wind"`
	Clouds     clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        sys       `json:"sys"`
	TimeZone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

func GetTemperature(w http.ResponseWriter, r *http.Request) {
	var data openweather
	apiKey := os.Getenv("OpenWeatherApiKey")

	response, _ := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, apiKey))
	bodyBytes, _ := io.ReadAll(response.Body)
	err := json.Unmarshal([]byte(string(bodyBytes)), &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%.2f", data.MainStats.Temp-kelvin)
}
