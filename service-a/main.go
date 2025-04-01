package main

import (
	"log"
	"net/http"

	"github.com/paulagates/cep-weather-tracing/service-a/handlers"
)

func main() {
	http.HandleFunc("/cep", handlers.HandleCEP)
	log.Println("Serviço A rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
