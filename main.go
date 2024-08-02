package main

import (
    "context"
    "fmt"
    "os"
    "time"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func initTracer() func() {
    fmt.Println("Initializing tracer...")
    jaegerEndpoint := os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT")
    fmt.Printf("Jaeger Endpoint: %s\n", jaegerEndpoint)

    exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)))
    if err != nil {
        fmt.Printf("Failed to create Jaeger exporter: %v\n", err)
        os.Exit(1)
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceName("github-actions-pipeline"),
        )),
    )
    otel.SetTracerProvider(tp)
    fmt.Println("Tracer initialized successfully")

    return func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            fmt.Printf("Error shutting down tracer provider: %v\n", err)
        }
    }
}

func main() {
    shutdown := initTracer()
    defer shutdown()

    tracer := otel.Tracer("github-actions-pipeline")
    ctx, span := tracer.Start(context.Background(), "main")
    defer span.End()

    fmt.Println("Starting pipeline simulation")

    // Simulating pipeline steps
    steps := []string{"Checkout code", "Set up Go", "Install dependencies", "Build application", "Run application"}
    for _, step := range steps {
        _, stepSpan := tracer.Start(ctx, step)
        fmt.Printf("Executing step: %s\n", step)
        time.Sleep(time.Second) // Simulate some work
        stepSpan.End()
        fmt.Printf("Completed step: %s\n", step)
    }

    fmt.Println("Pipeline simulation completed")
}
