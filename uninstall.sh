#!/bin/bash
# ─── CYPH3R AUTOMATED UNINSTALL ENGINE ────────────────────────────────

Reset="\033[0m"
Red="\033[31m"
NeonPink="\033[38;5;198m"

echo -e "${NeonPink}[*] Initializing CYPH3R Removal Sequence...${Reset}"

# Remove the system global binary pathway mappings
if [ -f /usr/local/bin/cyph3r ]; then
    echo -e "${Red}[-] Deleting global system executable path...${Reset}"
    sudo rm -f /usr/local/bin/cyph3r 2>/dev/null || rm -f /usr/local/bin/cyph3r 2>/dev/null
fi

# Clean localized folder directory parameters
if [ -f ./cyph3r ]; then
    echo -e "${Red}[-] Removing local compilation binaries...${Reset}"
    rm -f ./cyph3r
fi

echo -e "\033[38;5;82m[✓] CYPH3R has been cleanly uninstalled from your machine.${Reset}\n"
