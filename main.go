package main

import (
    "context"
    "log"
    "os"
    "time"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func initTracer() func() {
    log.Println("Initializing tracer...")
    jaegerEndpoint := os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT")
    log.Printf("Jaeger Endpoint: %s", jaegerEndpoint)

    exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)))
    if err != nil {
        log.Fatalf("Failed to create Jaeger exporter: %v", err)
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceName("github-actions-pipeline"),
        )),
    )
    otel.SetTracerProvider(tp)
    log.Println("Tracer initialized successfully")

    return func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            log.Printf("Error shutting down tracer provider: %v", err)
        }
    }
}

func main() {
    shutdown := initTracer()
    defer shutdown()

    tracer := otel.Tracer("github-actions-pipeline")
    ctx, span := tracer.Start(context.Background(), "main")
    defer span.End()

    log.Println("Starting pipeline simulation")

    // Simulating pipeline steps
    steps := []string{"Checkout code", "Set up Go", "Install dependencies", "Build application", "Run application"}
    for _, step := range steps {
        _, stepSpan := tracer.Start(ctx, step)
        log.Printf("Executing step: %s", step)
        time.Sleep(time.Second) // Simulate some work
        stepSpan.End()
        log.Printf("Completed step: %s", step)
    }

    log.Println("Pipeline simulation completed")
}