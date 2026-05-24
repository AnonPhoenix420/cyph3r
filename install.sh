#!/bin/bash

# ANSI Color Matrix for Installer HUD
Reset="\033[0m"
NeonBlue="\033;5;39m"
NeonGreen="\033[38;5;82m"
NeonPink="\033[38;5;198m"
Red="\033[31m"

echo -e "${NeonBlue}[*] Initializing CYPH3R Deployment Engine...${Reset}"

# 1. Verify Go compiler is present in the environment
if ! command -v go &> /dev/null; then
    echo -e "${Red}[-] Error: Go compiler not found. Install it via package manager first.${Reset}"
    exit 1
fi

# 2. Run the modern package layout build command
echo -e "${NeonBlue}[*] Compiling core dependencies and layout modules...${Reset}"
go build -o cyph3r ./cmd/cyph3r

if [ $? -ne 0 ]; then
    echo -e "${Red}[-] Compilation Failed. Check package path mapping codes.${Reset}"
    exit 1
fi

# 3. Grant local execution permissions
chmod +x cyph3r

# 4. Attempt to move binary to global system path for seamless execution
echo -e "${NeonBlue}[*] Copying binary matrix to global system path (/usr/local/bin/)...${Reset}"
if [ "$EUID" -ne 0 ]; then
    # If not running as root, attempt sudo
    sudo cp cyph3r /usr/local/bin/ 2>/dev/null
else
    cp cyph3r /usr/local/bin/ 2>/dev/null
fi

# Verification fallbacks specifically for Termux environments
if [ $? -eq 0 ]; then
    echo -e "${NeonGreen}[✓] CYPH3R installed globally! You can now run 'cyph3r' from anywhere.${Reset}"
else
    echo -e "${NeonPink}[!] Notice: Local deployment only. Run via standard local path: ./cyph3r${Reset}"
fi

echo -e "${NeonGreen}[+] CYPH3R installation cycle complete.${Reset}\n"
