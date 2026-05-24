# ─── CYPH3R v2.6 SYSTEM MAINTENANCE MAKEFILE ──────────────────────────

BINARY_NAME=cyph3r

all: build

build:
	@echo "[*] Building CYPH3R v2.6 production binary..."
	go build -o $(BINARY_NAME) ./cmd/cyph3r

repair:
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

.PHONY: all build repair clean
