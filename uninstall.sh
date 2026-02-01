#!/bin/bash

# CYPH3R v2.6 // Decommissioning Engine
# --------------------------------------
RED='\033[0;31m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}──[ CYPH3R DECOMMISSIONING ]──${NC}"

# 1. Binary Removal
if [ -f "./cyph3r" ]; then
    echo -e "${YELLOW}[*] Removing CYPH3R binary...${NC}"
    rm -f ./cyph3r
else
    echo -e "[!] Binary already removed or not found."
fi

# 2. Dependency & Cache Purge (Optional but recommended)
echo -e "${YELLOW}[*] Purging Go build cache...${NC}"
go clean -cache -modcache

# 3. Cleanup Build Artifacts
if [ -d "./bin" ]; then
    echo -e "${YELLOW}[*] Cleaning artifact directory...${NC}"
    rm -rf ./bin
fi

# 4. Final Verification
echo -e "\n${RED}[✔] SYSTEM WIPED SUCCESSFUL${NC}"
echo -e "${CYAN}Note: Your source code and go.mod remain intact.${NC}"
