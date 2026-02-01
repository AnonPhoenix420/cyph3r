#!/bin/bash

# CYPH3R v2.6 // Go 1.24 Deployment Engine
CYAN='\033[0;36m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${CYAN}──[ CYPH3R v2.6 INSTALLATION (Go 1.24) ]──${NC}"

# Verify Go 1.24+
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
if [[ $(echo -e "$GO_VERSION\n1.24" | sort -V | head -n1) != "1.24" ]]; then
    echo -e "${RED}[!] Error: This tool requires Go 1.24 or higher.${NC}"
    echo -e "${RED}[!] Current version: $GO_VERSION${NC}"
    exit 1
fi

echo -e "${GREEN}[*] Go 1.24 Runtime Detected.${NC}"

# Clean up any old module mismatches
rm -f go.sum
go mod tidy

# Build
echo -e "${CYAN}[*] Compiling CYPH3R...${NC}"
go build -o cyph3r ./cmd/cyph3r

if [ -f "./cyph3r" ]; then
    chmod +x cyph3r
    echo -e "${GREEN}[✔] System Optimized for Go 1.24 and Deployed.${NC}"
else
    echo -e "${RED}[!] Build failed.${NC}"
    exit 1
fi
