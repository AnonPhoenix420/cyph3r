#!/bin/bash
echo "[*] Initializing CYPH3R Installation..."

# 1. Tidy modules and build
go mod tidy
go build -o cyph3r ./cmd/cyph3r/main.go

if [ $? -eq 0 ]; then
    echo "[+] Build Successful."
else
    echo "[-] Build Failed. Check your Go environment."
    exit 1
fi

# 2. Move to system path
sudo mv cyph3r /usr/local/bin/

# 3. Final verification
if command -v cyph3r &> /dev/null; then
    echo "[+] CYPH3R installed to /usr/local/bin/cyph3r"
    echo "[!] Run it using: cyph3r -target google.com"
else
    echo "[-] Installation failed to map to PATH."
fi
