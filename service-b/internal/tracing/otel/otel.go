package otelconfig

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func InitProvider(serviceName, zipkinURL string) (func(context.Context) error, error) {
	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure(),                      // Sem TLS (para ambiente local)
		otlptracegrpc.WithEndpoint("otel-collector:4317"), // Endpoint do Collector
	)

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar recurso: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	log.Println("✅ OpenTelemetry configurado para Zipkin em:", zipkinURL)

	return tracerProvider.Shutdown, nil
}
