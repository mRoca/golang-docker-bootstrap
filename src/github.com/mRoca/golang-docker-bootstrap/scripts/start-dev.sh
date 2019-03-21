#!/bin/sh

while true; do
  go build -o bin/cli ./cmd/cli/main.go
  echo "bin/cli built with success"
  inotifywait -r -e modify cmd internal scripts
done
