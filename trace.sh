#!/bin/bash

# Download OpenTelemetry CLI
if [ ! -f opentelemetry-instrument ]; then
  curl -L https://github.com/open-telemetry/opentelemetry-go-instrumentation/releases/latest/download/opentelemetry-instrument-linux-amd64 -o opentelemetry-instrument
  chmod +x opentelemetry-instrument
fi

# Set up OpenTelemetry exporter
export OTEL_EXPORTER_JAEGER_ENDPOINT=${OTEL_EXPORTER_JAEGER_ENDPOINT}
export OTEL_SERVICE_NAME=github-actions-pipeline

# Run the command with OpenTelemetry instrumentation
./opentelemetry-instrument run "$@"
