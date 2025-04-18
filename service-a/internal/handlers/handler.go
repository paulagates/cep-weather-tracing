package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/paulagates/cep-weather-tracing/service-a/internal/services"
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

	var reqBody services.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isValidCEP := func(cep string) bool {
		match, _ := regexp.MatchString(`^\d{8}$`, cep)
		return match
	}

	if !isValidCEP(reqBody.CEP) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	res, err := services.ForwardToServiceB(reqBody)
	if err != nil {
		log.Println("Failed to forward to Service B: ", err)
		http.Error(w, "Failed to forward to Service B", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to read response body: ", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	w.Write(body)

}
