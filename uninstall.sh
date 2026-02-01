
#!/bin/bash
# CYPH3R v2.6 // Final Cleanup Script

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${RED}[!] Initiating CYPH3R Decommissioning...${NC}"

# 1. Remove the binary
if [ -f "./cyph3r" ]; then
    rm ./cyph3r
    echo -e "${GREEN}[✔] Binary removed.${NC}"
fi

# 2. Clean Go Cache (Optional but recommended)
echo -e "[*] Purging module cache..."
go clean -modcache

echo -e "${GREEN}[✔] Uninstall Complete. System Cleaned.${NC}"

