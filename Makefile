# ─── CYPH3R v2.6 SYSTEM MAINTENANCE MAKEFILE ──────────────────────────

BINARY_NAME=cyph3r

all: build

fix-imports:
	@echo "[*] Normalizing module import paths..."
	@find . -type f -name '*.go' -exec sed -i 's|"cyph3r/internal|"github.com/AnonPhoenix420/cyph3r/internal|g' {} +

build: fix-imports
	@echo "[*] Syncing dependencies..."
	@go mod tidy
	@echo "[*] Building CYPH3R v2.6 production binary..."
	go build -o $(BINARY_NAME) ./cmd/cyph3r
	@echo "[✓] Build complete."

repair: fix-imports
	@echo "[*] Initializing CYPH3R Self-Repair Routine..."
	go clean -modcache
	go mod tidy
	go build -o $(BINARY_NAME) ./cmd/cyph3r
	@echo "[✓] Environment successfully repaired and resynced."

clean:
	@echo "[*] Removing build artifacts and binaries..."
	rm -f $(BINARY_NAME)
	go clean
	@echo "[✓] Project environment cleaned."

.PHONY: all build repair clean fix-imports
