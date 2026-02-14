# CYPH3R Tactical Makefile
BINARY_NAME=cyph3r
BUILD_DIR=./bin
MAIN_PATH=./cmd/cyph3r/main.go

.PHONY: all build clean install repair docker backup help

all: repair build

## build: Compiles the binary to the bin directory
build:
	@echo "[*] Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "[+] Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

## clean: Removes build artifacts
clean:
	@echo "[*] Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "[+] Workspace clean."

## repair: Forces a module refresh and fixes dependencies
repair:
	@echo "[!] Initializing Toolchain Repair..."
	@go mod tidy
	@go mod verify
	@echo "[+] Dependencies verified."

## install: Builds and moves the binary to /usr/local/bin
install: build
	@echo "[*] Installing to system path..."
	@sudo mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "[+] Installation successful. Type 'cyph3r' to run."

## docker: Builds the Docker image
docker:
	@docker build -t $(BINARY_NAME):latest .

## backup: Runs the backup script
backup:
	@chmod +x backup.sh
	@./backup.sh

help:
	@echo "CYPH3R Build System Commands:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
