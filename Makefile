BINARY_NAME=cyph3r

all: build

build:
	@echo "[*] Syncing dependencies..."
	go mod tidy
	@echo "[*] Building CYPH3R v2.6 production binary..."
	go build -o $(BINARY_NAME) ./cmd/cyph3r
	@echo "[✓] Build complete."

clean:
	@echo "[*] Removing build artifacts..."
	rm -f $(BINARY_NAME)
	go clean
	@echo "[✓] Environment cleaned."

.PHONY: all build clean
