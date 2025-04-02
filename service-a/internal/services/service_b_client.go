package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type RequestBody struct {
	CEP string `json:"cep"`
}

func ForwardToServiceB(reqBody RequestBody) (*http.Response, error) {
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("SERVICE_B_URL"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
