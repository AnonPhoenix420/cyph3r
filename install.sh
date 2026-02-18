#!/bin/bash
# Cyph3r Standard Installer

echo -e "\033[38;5;39m[*] Setting up Cyph3r Environment...\033[0m"

# Set permissions for scripts
chmod +x backup.sh uninstall.sh

# Build and Install
make install

if [ $? -eq 0 ]; then
    echo -e "\033[38;5;82m[+] Installation Complete. Execute 'cyph3r -t <target>' to begin.${NC}"
else
    echo -e "\033[31m[!] Installation failed. Ensure Go is installed.${NC}"
fi
