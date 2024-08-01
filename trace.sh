#!/bin/bash

# Install OpenTelemetry CLI
curl -L https://github.com/open-telemetry/opentelemetry-go-instrumentation/releases/latest/download/opentelemetry-instrument-linux-amd64 -o opentelemetry-instrument
chmod +x opentelemetry-instrument

# Set up OpenTelemetry exporter
export OTEL_EXPORTER_JAEGER_ENDPOINT=http://<ngrok-url>/api/traces
export OTEL_SERVICE_NAME=github-actions-pipeline

# Run the command with OpenTelemetry instrumentation
./opentelemetry-instrument run "$@"
