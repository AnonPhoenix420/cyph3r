#!/usr/bin/env bash
# ============================================================
# CYPH3R Installation Script
# Fully hands-free: clones, builds, and prepares the tool
# Works on: Parrot OS, Linux, Termux (Android)
# ============================================================

set -e

# ================= CONFIG =================
REPO_URL="https://github.com/AnonPhoenix420/cyph3r.git"
BINARY_NAME="cyph3r"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"

# ================= CHECK ENV =================
if ! command -v go &>/dev/null; then
    echo "[INFO] Go not found. Installing Go..."
    OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH_TYPE=$(uname -m)

    if [[ "$OS_TYPE" == "linux" ]]; then
        if [[ "$ARCH_TYPE" == "aarch64" ]]; then
            GO_TAR="go1.23.0.linux-arm64.tar.gz"
        else
            GO_TAR="go1.23.0.linux-amd64.tar.gz"
        fi
        curl -LO https://go.dev/dl/$GO_TAR
        sudo tar -C /usr/local -xzf $GO_TAR
        export PATH=$PATH:/usr/local/go/bin
    else
        echo "[ERROR] Unsupported OS: $OS_TYPE"
        exit 1
    fi
fi

# ================= CLONE REPO =================
if [[ ! -d "cyph3r" ]]; then
    echo "[INFO] Cloning CYPH3R repository..."
    git clone "$REPO_URL"
fi

cd cyph3r

# ================= BUILD =================
echo "[INFO] Initializing Go modules..."
go mod tidy

echo "[INFO] Building $BINARY_NAME..."
go build -o "$BINARY_NAME" main.go intel.go

# ================= INSTALL =================
if [[ "$EUID" -ne 0 && "$(uname -s)" != "Linux" ]]; then
    echo "[INFO] Build complete. Run ./cyph3r to start."
else
    echo "[INFO] Installing system-wide..."
    sudo install -m 755 "$BINARY_NAME" "$INSTALL_PATH"
    echo "[INFO] Installed as $INSTALL_PATH"
fi

echo "[SUCCESS] CYPH3R is ready!"
