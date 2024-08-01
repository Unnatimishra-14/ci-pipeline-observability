#!/bin/bash

# Install OpenTelemetry CLI
curl -L https://github.com/open-telemetry/opentelemetry-go-instrumentation/releases/latest/download/opentelemetry-instrument-linux-amd64 -o opentelemetry-instrument
chmod +x opentelemetry-instrument

# Set up OpenTelemetry exporter
export OTEL_EXPORTER_JAEGER_ENDPOINT=http://https://ff9b-2401-4900-1c83-b5c4-48f0-79a2-9814-9a0c.ngrok-free.app/search/api/traces
export OTEL_SERVICE_NAME=github-actions-pipeline

# Run the command with OpenTelemetry instrumentation
./opentelemetry-instrument run "$@"
