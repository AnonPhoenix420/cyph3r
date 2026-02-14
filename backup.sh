#!/bin/bash

# Configuration
BACKUP_DIR="./backups"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_NAME="cyph3r_backup_$TIMESTAMP.tar.gz"

# 1. Create backup directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

# 2. Compress relevant folders (excluding large build artifacts/backups)
echo "[*] Creating backup: $BACKUP_NAME"
tar -czf "$BACKUP_DIR/$BACKUP_NAME" \
    --exclude="./backups" \
    --exclude="./cyph3r" \
    ./cmd ./internal go.mod

if [ $? -eq 0 ]; then
    echo "[+] Backup saved to: $BACKUP_DIR/$BACKUP_NAME"
else
    echo "[-] Backup failed."
fi
