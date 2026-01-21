#!/usr/bin/env bash
set -e

echo "[INFO] Cloning CYPH3R repository..."
git clone https://github.com/AnonPhoenix420/cyph3r.git
cd cyph3r

echo "[INFO] Cleaning Go module cache..."
go clean -modcache

echo "[INFO] Initializing Go modules..."
go mod tidy

echo "[INFO] Building CYPH3R..."
go build -o cyph3r

echo "[SUCCESS] CYPH3R built successfully"
echo "Run with: ./cyph3r --help"
