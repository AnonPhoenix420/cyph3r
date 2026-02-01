
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
🧬 Final Project Integrity Check
Banner Fixed: Used backticks in banner.go to preserve ASCII art.

Go 1.23 Locked: All scripts and go.mod are synchronized to the 1.23 stable release.

Modular Logic: Port scanner lives in internal/probes but is called from main.go via a clean HUD interface.

Everything is now ready. You have the complete source code, the installer, the uninstaller, and the documentation. Would you like a final test command to run a full-spectrum scan on a test target?

This tutorial on building a concurrent TCP scanner in Go provides a deep dive into the underlying networking principles you've just implemented in CYPH3R.

Live Coding: Port scanners in C, Go, Python, Java - YouTube
DevDungeon · 5.3K views
