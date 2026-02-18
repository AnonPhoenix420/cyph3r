# Cyph3r Tactical Build System
BINARY_NAME=cyph3r
INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)

all: build

build:
	@echo "[\033[38;5;99m*\033[0m] Compiling Cyph3r God Mode..."
	@go build -o $(BINARY_NAME) ./cmd/cyph3r
	@chmod +x $(BINARY_NAME)

install: build
	@echo "[\033[38;5;82m+\033[0m] Installing to /usr/local/bin..."
	@sudo cp $(BINARY_NAME) /usr/local/bin/
	@echo "[\033[38;5;82m+\033[0m] Success. You can now run 'cyph3r' from any directory."

uninstall:
	@echo "[\033[31m!\033[0m] Removing Cyph3r from system path..."
	@if [ -f $(INSTALL_PATH) ]; then \
		sudo rm -f $(INSTALL_PATH); \
		echo "[\033[38;5;82m+\033[0m] System binary removed."; \
	else \
		echo "[\033[31m!\033[0m] Binary not found in /usr/local/bin."; \
	fi
	@rm -f $(BINARY_NAME)
	@echo "[\033[38;5;198m*\033[0m] Cleanup complete."

clean:
	@go clean
	@rm -f $(BINARY_NAME)
	@rm -rf backups/
