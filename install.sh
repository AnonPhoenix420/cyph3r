#!/bin/bash

# CYPH3R v2.6 // Deployment Engine
# High-Density Intelligence HUD Installer

# Text Formatting
CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${CYAN}──[ CYPH3R v2.6 INSTALLATION ]──${NC}"

# 1. Verification: Go Environment
echo -ne "${WHITE}[*] Verifying Go Runtime... "
if ! command -v go &> /dev/null; then
    echo -e "${RED}[FAILED]${NC}"
    echo -e "${YELLOW}[!] Error: Go 1.23+ is required for Node Intelligence features.${NC}"
    exit 1
fi
GO_VER=$(go version | awk '{print $3}' | sed 's/go//')
echo -e "${GREEN}[READY] (v$GO_VER)${NC}"

# 2. Dependency Injection
echo -e "${CYAN}[*] Synchronizing Advanced Intel Modules...${NC}"
# Ensuring the new whois-go and color libraries are present
go mod tidy
go get github.com/likexian/whois-go
go get github.com/fatih/color

# 3. Build Sequence
echo -e "${CYAN}[*] Compiling High-Density Binary...${NC}"
if [ -f "Makefile" ]; then
    make repair
    make build
else
    # Fallback if Makefile is missing
    go build -o cyph3r ./cmd/cyph3r
fi

# 4. Finalization & Permissions
if [ -f "./cyph3r" ]; then
    chmod +x cyph3r
    echo -e "\n${GREEN}[✔] INSTALLATION COMPLETE${NC}"
    echo -e "${WHITE}───────────────────────────────────────"
    echo -e "${CYAN}Launch HUD:${NC}  ./cyph3r --target <host>"
    echo -e "${CYAN}Live Feed: ${NC}  ./cyph3r --target <host> --monitor"
    echo -e "${WHITE}───────────────────────────────────────${NC}"
else
    echo -e "${RED}[!] Critical: Build failed. Check compiler output above.${NC}"
    exit 1
fi
