#!/bin/bash

# --- Color Definitions ---
BLUE='\033[38;5;33m'
GREEN='\033[38;5;82m'
PINK='\033[38;5;201m'
RESET='\033[0m'

INSTALL_PATH="/usr/local/bin/cyph3r"

echo -e "${PINK}──────────────────────────────────────────────────${RESET}"
echo -e "${PINK}         CYPH3R: SYSTEM DECOMMISSION             ${RESET}"
echo -e "${PINK}──────────────────────────────────────────────────${RESET}"

# 1. Check if binary exists in the system path
if [ -f "$INSTALL_PATH" ]; then
    echo -e "${BLUE}[*] Found system binary at $INSTALL_PATH${RESET}"
    echo -e "${BLUE}[*] Requesting sudo permissions for removal...${RESET}"
    
    sudo rm "$INSTALL_PATH"
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}[+] SUCCESS: Binary removed from system path.${RESET}"
    else
        echo -e "${PINK}[!] ERROR: Failed to remove system binary.${RESET}"
        exit 1
    fi
else
    echo -e "${BLUE}[*] No system binary found in $INSTALL_PATH. Skipping...${RESET}"
fi

# 2. Local Cleanup (Optional but recommended)
echo -e "${BLUE}[*] Cleaning local build artifacts and Go cache...${RESET}"
rm -f cyph3r
go clean -cache

echo -e "${GREEN}[+] DECOMMISSION COMPLETE.${RESET}"
echo -e "${PINK}──────────────────────────────────────────────────${RESET}"
