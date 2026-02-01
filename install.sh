#!/bin/bash
# CYPH3R v2.6 // Android Deployment
echo -e "\033[0;36m──[ CYPH3R ANDROID BUILD ]──\033[0m"

# Ensure Go is present
if ! command -v go &> /dev/null; then
    echo "Installing Go..."
    pkg install golang -y
fi

# Clean & Build
rm -f go.sum
go mod tidy
go build -o cyph3r ./cmd/cyph3r

if [ -f "./cyph3r" ]; then
    chmod +x cyph3r
    echo -e "\033[0;32m[✔] CYPH3R HUD Ready. Run with ./cyph3r\033[0m"
else
    echo -e "\033[0;31m[!] Build failed. Check error logs.\033[0m"
fi
