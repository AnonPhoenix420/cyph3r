#!/bin/bash

# --- Color Definitions ---
BLUE='\033[38;5;33m'
GREEN='\033[38;5;82m'
PINK='\033[38;5;201m'
RESET='\033[0m'

echo -e "${BLUE}──────────────────────────────────────────────────${RESET}"
echo -e "${BLUE}         CYPH3R: NEON TECH INSTALLER             ${RESET}"
echo -e "${BLUE}──────────────────────────────────────────────────${RESET}"

# 1. Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${PINK}[!] ERROR: Go is not installed. Please install Go 1.23+ and try again.${RESET}"
    exit 1
fi

# 2. Sync dependencies
echo -e "${BLUE}[*] Synchronizing dependencies...${RESET}"
go mod tidy
if [ $? -ne 0 ]; then
    echo -e "${PINK}[!] ERROR: Failed to sync modules.${RESET}"
    exit 1
fi

# 3. Build the binary
echo -e "${BLUE}[*] Compiling binary for host architecture...${RESET}"
go build -o cyph3r ./cmd/cyph3r
if [ $? -ne 0 ]; then
    echo -e "${PINK}[!] ERROR: Compilation failed.${RESET}"
    exit 1
fi

# 4. Install to system path
echo -e "${BLUE}[*] Requesting sudo permissions to move binary to /usr/local/bin...${RESET}"
sudo mv cyph3r /usr/local/bin/cyph3r
sudo chmod +x /usr/local/bin/cyph3r

# 5. Final Verification
if command -v cyph3r &> /dev/null; then
    echo -e "${GREEN}[+] SUCCESS: Cyph3r is installed and ready.${RESET}"
    echo -e "${GREEN}[+] Usage: cyph3r --target <domain> --scan${RESET}"
else
    echo -e "${PINK}[!] ERROR: Installation failed during path move.${RESET}"
    exit 1
fi

echo -e "${BLUE}──────────────────────────────────────────────────${RESET}"
