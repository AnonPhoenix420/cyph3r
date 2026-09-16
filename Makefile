BINARY_NAME=cyph3r

all: build

build:
	@echo "[*] Syncing dependencies..."
	go mod tidy
	@echo "[*] Building CYPH3R v2.6 production binary..."
	go build -o $(BINARY_NAME) ./cmd/cyph3r
	@echo "[✓] Build complete."

cross-compile:
	@echo "[*] Cross-compiling multi-platform binaries..."
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)_linux_amd64 ./cmd/cyph3r
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)_windows_amd64.exe ./cmd/cyph3r
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)_darwin_amd64 ./cmd/cyph3r
	@echo "[✓] Cross-compilation complete."

clean:
	@echo "[*] Removing build artifacts..."
	rm -f $(BINARY_NAME)
	rm -rf dist/
	go clean
	@echo "[✓] Environment cleaned."

.PHONY: all build cross-compile clean
