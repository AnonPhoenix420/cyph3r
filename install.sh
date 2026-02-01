#!/bin/bash

# CYPH3R v2.6 - HUD Automated Installer
# ------------------------------------
CYAN='\033[0;36m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${CYAN}[*] Initializing CYPH3R HUD Installation...${NC}"

# 1. Check for Go
if ! command -v go &> /dev/null; then
    echo -e "${RED}[!] Go 1.23+ not found. Please install Go first.${NC}"
    exit 1
fi

# 2. Sync Dependencies
echo -e "${CYAN}[*] Syncing HUD Modules...${NC}"
go mod tidy

# 3. Build Binary
echo -e "${CYAN}[*] Compiling CYPH3R v2.6...${NC}"
go build -o cyph3r ./cmd/cyph3r

if [ $? -eq 0 ]; then
    echo -e "${GREEN}[✔] Installation Complete!${NC}"
    echo -e "${GREEN}[✔] Run with: ./cyph3r --target google.com${NC}"
else
    echo -e "${RED}[!] Build failed. Check your Go version (requires 1.23+).${NC}"
fi
