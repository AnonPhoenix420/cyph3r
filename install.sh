#!/bin/bash
# Go 1.23 Environment
rm -f go.sum
go mod tidy
go build -o cyph3r ./cmd/cyph3r
chmod +x cyph3r
echo "CYPH3R v2.6 Ready."
