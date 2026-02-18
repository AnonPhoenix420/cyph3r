#!/bin/bash
echo -e "\033[38;5;39m[*] Setting up Cyph3r Environment...\033[0m"

# Ensure script permissions
chmod +x backup.sh

# Run Go Tidy to verify dependencies
go mod tidy

# Build the tool
make build

if [ -f "./cyph3r" ]; then
    echo -e "\033[38;5;82m[+] Cyph3r is ready. Run with ./cyph3r -t <target>\033[0m"
else
    echo -e "\033[31m[!] Build failed. Check Go installation.\033[0m"
fi
