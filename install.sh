#!/bin/bash
# CYPH3R v2.6 Automated Installer

CYAN='\033[0;36m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${CYAN}──[ CYPH3R INSTALLATION START ]──${NC}"

# Step 1: Check Go Version
if ! command -v go &> /dev/null; then
    echo -e "${RED}[!] Go is not installed. Please follow the README instructions.${NC}"
    exit 1
fi

# Step 2: Architecture Check
ARCH=$(uname -m)
echo -e "[*] System Architecture: ${CYAN}$ARCH${NC}"

# Step 3: Dependency Sync
echo -e "[*] Syncing modules..."
go mod tidy

# Step 4: Build
echo -e "[*] Compiling CYPH3R..."
go build -o cyph3r ./cmd/cyph3r

if [ -f "./cyph3r" ]; then
    chmod +x cyph3r
    echo -e "${GREEN}[✔] Installation Successful.${NC}"
    echo -e "Usage: ./cyph3r --target google.com --scan"
else
    echo -e "${RED}[!] Build failed. Try 'make repair'.${NC}"
fi
