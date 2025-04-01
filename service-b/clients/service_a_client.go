package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type City struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func GetCityFromCEP(cep string) (string, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("can not find zipcode")
	}

	var city City
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(body, &city); err != nil {
		return "", err
	}

	return city.Localidade, nil
}

type Temperature struct {
	TempC float64 `json:"temp_C"`
}

// WeatherResponse define o formato da resposta da API de clima
type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

// GetTemperatureFromCity consulta a API de clima para obter a temperatura da cidade
func GetTemperatureFromCity(city string) (WeatherResponse, error) {
	apiKey := "83341080b0c04ea49a1140041250603"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, errors.New("failed to fetch weather data")
	}

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
