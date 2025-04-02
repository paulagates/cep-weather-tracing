package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/paulagates/cep-weather-tracing/service-b/internal/services"
	"github.com/paulagates/cep-weather-tracing/service-b/internal/utils"
)

type Response struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func HandleCEP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cep, exists := reqBody["cep"]
	if !exists {
		http.Error(w, "CEP not provided", http.StatusBadRequest)
		return
	}

	isValidCEP := func(cep string) bool {
		match, _ := regexp.MatchString(`^\d{8}$`, cep)
		return match
	}

	if !isValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	city, err := services.GetCityFromCEP(cep)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	temp, err := services.GetTemperatureFromCity(city)
	if err != nil {
		log.Println("Failed to get temperature: ", err)
		http.Error(w, "Failed to get temperature", http.StatusInternalServerError)
		return
	}

	tempC := temp.Current.TempC
	tempF := utils.ConvertCelsiusToFahrenheit(tempC)
	tempK := utils.ConvertCelsiusToKelvin(tempC)

	response := Response{
		City:  city,
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
