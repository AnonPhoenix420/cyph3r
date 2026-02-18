#!/bin/bash
# Cyph3r Emergency Uninstall Script

RED='\033[31m'
GREEN='\033[38;5;82m'
PURPLE='\033[38;5;99m'
NC='\033[0m'

echo -e "${PURPLE}[*] Initializing Cyph3r Uninstallation...${NC}"

# 1. Remove from local directory
if [ -f "./cyph3r" ]; then
    rm -f ./cyph3r
    echo -e "${GREEN}[+] Local binary removed.${NC}"
fi

# 2. Remove from System Path
if [ -f "/usr/local/bin/cyph3r" ]; then
    sudo rm -f /usr/local/bin/cyph3r
    echo -e "${GREEN}[+] System-wide binary removed.${NC}"
fi

# 3. Clean up temporary logs/backups (Optional - uncomment if desired)
# rm -rf ./backups
# rm -f *.log

echo -e "${PURPLE}[*] Cyph3r has been scrubbed from this node.${NC}"
