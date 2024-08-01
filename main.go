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
    "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer() func() {
    exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT"))))
    if err != nil {
        fmt.Printf("Failed to create Jaeger exporter: %v\n", err)
        return func() {}
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("github-actions-pipeline"),
        )),
    )

    otel.SetTracerProvider(tp)

    return func() {
        _ = tp.Shutdown(context.Background())
    }
}

func main() {
    shutdown := initTracer()
    defer shutdown()

    tracer := otel.Tracer("example-tracer")
    _, span := tracer.Start(context.Background(), "main")
    defer span.End()

    fmt.Println("Starting the application...")
    time.Sleep(5 * time.Second)
    fmt.Println("Application finished.")
}