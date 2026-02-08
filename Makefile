# Project Variables
BINARY_NAME=cyph3r
SOURCE_PATH=./cmd/cyph3r
INSTALL_PATH=/usr/local/bin

# Colors for terminal output
BLUE=\033[38;5;33m
GREEN=\033[38;5;82m
PINK=\033[38;5;201m
RESET=\033[0m

.PHONY: all build sync install clean help

all: sync build

## sync: Download dependencies and tidy go.mod
sync:
	@echo "$(BLUE)[*] Syncing Neon Tech Dependencies...$(RESET)"
	go mod tidy

## build: Compile the binary
build:
	@echo "$(BLUE)[*] Compiling $(BINARY_NAME)...$(RESET)"
	go build -o $(BINARY_NAME) $(SOURCE_PATH)
	@echo "$(GREEN)[+] Build Complete: ./$(BINARY_NAME)$(RESET)"

## install: Move binary to /usr/local/bin (Requires sudo)
install: build
	@echo "$(BLUE)[*] Installing $(BINARY_NAME) to $(INSTALL_PATH)...$(RESET)"
	@sudo cp $(BINARY_NAME) $(INSTALL_PATH)
	@sudo chmod +x $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "$(GREEN)[+] Installation Successful. You can now run 'cyph3r' from anywhere.$(RESET)"

## clean: Remove binary and cached files
clean:
	@echo "$(PINK)[!] Cleaning workspace...$(RESET)"
	rm -f $(BINARY_NAME)
	go clean -cache
	@echo "$(GREEN)[+] Workspace cleared.$(RESET)"

## help: Display available commands
help:
	@echo "$(BLUE)Cyph3r Build System - Commands:$(RESET)"
	@echo "  $(GREEN)make sync$(RESET)    - Tidy go.mod and download libraries"
	@echo "  $(GREEN)make build$(RESET)   - Build the cyph3r binary"
	@echo "  $(GREEN)make install$(RESET) - Build and install to system path"
	@echo "  $(GREEN)make clean$(RESET)   - Remove binary and build cache"
