package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/paulagates/cep-weather-tracing/service-a/internal/handlers"
	otelconfig "github.com/paulagates/cep-weather-tracing/service-a/internal/tracing/otel"
	"go.opentelemetry.io/otel"
)

func main() {
	godotenv.Load("../.env")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	shutdown, err := otelconfig.InitProvider(os.Getenv("OTEL_SERVICE_NAME"), os.Getenv("OTEL_EXPORTER_ZIPKIN_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("Erro ao encerrar OpenTelemetry: %v", err)
		}
	}()
	tracer := otel.Tracer("service-a-tracer")

	ctx, span := tracer.Start(ctx, "OperacaoPrincipal")
	time.Sleep(500 * time.Millisecond)
	span.End()

	mux := http.NewServeMux()
	mux.HandleFunc("/cep", handlers.HandleCEP)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("🚀 Serviço rodando na porta 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("🛑 Encerrando serviço...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Erro ao encerrar servidor: %v", err)
	}

	log.Println("✅ Serviço encerrado com sucesso!")
}
