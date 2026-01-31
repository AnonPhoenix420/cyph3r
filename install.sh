#!/bin/bash
set -e

# Create folder structure
mkdir -p cmd/cyph3r internal/intel internal/output

# Initialize Go module if not exists
if [ ! -f go.mod ]; then
    go mod init github.com/AnonPhoenix420/cyph3r
fi

echo "[*] Cleaning environment..."
rm -f go.sum
go clean -modcache

echo "[*] Downloading dependencies..."
go get github.com/nyaruka/phonenumbers
go mod tidy

echo "[*] Compiling CYPH3R..."
go build -o cyph3r ./cmd/cyph3r

echo "[✔] Done! Run: ./cyph3r --target google.com --proto https"
