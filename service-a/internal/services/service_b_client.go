package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const serviceBURL = "http://localhost:8081/cep"

type RequestBody struct {
	CEP string `json:"cep"`
}

func ForwardToServiceB(reqBody RequestBody) (*http.Response, error) {
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(serviceBURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
