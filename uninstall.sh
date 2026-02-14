#!/bin/bash
echo "[*] Uninstalling CYPH3R..."

# Remove the binary from system path
if [ -f /usr/local/bin/cyph3r ]; then
    sudo rm /usr/local/bin/cyph3r
    echo "[+] Binary removed from /usr/local/bin/cyph3r"
else
    echo "[!] Binary not found in /usr/local/bin/cyph3r"
fi

echo "[+] Uninstallation Complete."
