#!/bin/bash

# CYPH3R v2.6 - HUD System Decommissioner
# ---------------------------------------
RED='\033[0;31m'
NC='\033[0m'
CYAN='\033[0;36m'

echo -e "${RED}[!] Initializing CYPH3R Decommissioning Sequence...${NC}"

# 1. Remove the compiled binary
if [ -f "./cyph3r" ]; then
    rm ./cyph3r
    echo -e "${CYAN}[*] Local binary removed.${NC}"
fi

# 2. Clean Go Mod cache for these specific dependencies
echo -e "${CYAN}[*] Purging build artifacts...${NC}"
go clean -cache
go clean -modcache

# 3. Optional: Remove the entire directory
echo -e "${RED}[?] Would you like to delete the entire 'cyph3r' source directory? (y/n)${NC}"
read -r choice

if [ "$choice" == "y" ] || [ "$choice" == "Y" ]; then
    cd ..
    rm -rf cyph3r
    echo -e "${RED}[✔] Repository purged. System clean.${NC}"
else
    echo -e "${CYAN}[*] Source files preserved. Binary and cache cleared.${NC}"
fi
