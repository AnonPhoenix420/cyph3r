#!/bin/bash
set -e

echo "[INFO] Cleaning old Go modules..."
go clean -modcache
rm -f go.sum

echo "[INFO] Initializing Go modules..."
go mod tidy

echo "[INFO] Building CYPH3R..."
go build -o cyph3r ./cmd/cyph3r

echo "[INFO] Build completed!"
echo "You can now run './cyph3r --help' to see usage."
