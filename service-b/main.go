package main

import (
	"log"
	"net/http"

	"github.com/paulagates/cep-weather-tracing/service-b/handlers"
)

func main() {
	http.HandleFunc("/cep", handlers.HandleCEP)
	log.Println("Serviço B rodando na porta 8081...")
	http.ListenAndServe(":8081", nil)
}
