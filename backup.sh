#!/bin/bash
# Cyph3r Source Persistence Script

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_DIR="./backups"

mkdir -p $BACKUP_DIR

echo -e "\033[38;5;198m[*] Initializing Source Archive...\033[0m"

# Archive the core logic and command structure
tar -czf "$BACKUP_DIR/cyph3r_backup_$TIMESTAMP.tar.gz" \
    cmd/ \
    internal/ \
    go.mod \
    go.sum \
    Makefile \
    Dockerfile

if [ $? -eq 0 ]; then
    echo -e "\033[38;5;82m[+] Backup Successful: $BACKUP_DIR/cyph3r_backup_$TIMESTAMP.tar.gz\033[0m"
else
    echo -e "\033[31m[!] Backup Failed!\033[0m"
fi
