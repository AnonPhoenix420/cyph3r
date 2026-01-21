#!/bin/bash
set -e

echo "[INFO] Installing dependencies..."

# Detect OS / architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
if [[ "$ARCH" == "aarch64" ]]; then ARCH="arm64"; fi

# Ensure git is installed
if ! command -v git &> /dev/null; then
    echo "[INFO] Installing git..."
    if [[ "$OS" == "linux" ]]; then
        sudo apt update && sudo apt install -y git
    else
        echo "[ERROR] Git not found. Install it manually."
        exit 1
    fi
fi

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo "[INFO] Installing Go..."
    GO_VER="1.23.0"
    curl -LO https://go.dev/dl/go${GO_VER}.${OS}-${ARCH}.tar.gz
    sudo tar -C /usr/local -xzf go${GO_VER}.${OS}-${ARCH}.tar.gz
    export PATH=$PATH:/usr/local/go/bin
fi

echo "[INFO] Cloning CYPH3R repository..."
git clone https://github.com/AnonPhoenix420/cyph3r.git
cd cyph3r

echo "[INFO] Adding go.mod and go.sum..."
# Drop-in go.mod
cat > go.mod <<'EOF'
module github.com/AnonPhoenix420/cyph3r

go 1.23

require (
    github.com/kr/text v0.2.0
    github.com/nyaruka/phonenumbers v1.1.0
    github.com/prometheus/client_golang v1.23.2
)

replace github.com/nyaruka/phonenumbers => github.com/nyaruka/phonenumbers v1.1.0
EOF

# Drop-in go.sum
cat > go.sum <<'EOF'
github.com/kr/text v0.2.0 h1:XixvHkcdZVEA5E3yP4dEIV1vZl1pY/SlmZr1mcF0roI=
github.com/kr/text v0.2.0/go.mod h1:1pK/2kckD8YcHbSEU6fzfY0VdhlHEg3y4N58mXU2wX4=
github.com/nyaruka/phonenumbers v1.1.0 h1:9kW3ZYnYjv1M8j6Qx7+1XhTSxkwj5s1j6nzE/1zp4Hw=
github.com/nyaruka/phonenumbers v1.1.0/go.mod h1:SkA5aI9/X5wmp+TeG+a0Dqyk+cLw5VksK09oedx3/BM=
github.com/prometheus/client_golang v1.23.2 h1:8FMRaV+ZwoZqUCM6q9q1KfFYpMf8ICq0+XIL2+HUBqU=
github.com/prometheus/client_golang v1.23.2/go.mod h1:KX/hkBR6fSx5Yt25/PuBMe2XsoL3z7CRUPOY8TtX5uw=
EOF

echo "[INFO] Initializing Go modules..."
go mod tidy

echo "[INFO] Building CYPH3R..."
go build -o cyph3r ./cmd/cyph3r

echo "[SUCCESS] CYPH3R built successfully!"
echo "Run it with: ./cyph3r --help"
