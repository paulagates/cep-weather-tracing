package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Temperature struct {
	TempC float64 `json:"temp_C"`
}

type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureFromCity(city string) (WeatherResponse, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, url.QueryEscape(city))
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()
	var weather WeatherResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	if err := json.Unmarshal(body, &weather); err != nil {
		return WeatherResponse{}, err
	}

	return weather, nil
}
