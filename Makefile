# Cyph3r Tactical Build System
BINARY_NAME=cyph3r

all: build

build:
	@echo "[\033[38;5;99m*\033[0m] Compiling Cyph3r God Mode..."
	@go build -o $(BINARY_NAME) ./cmd/cyph3r
	@chmod +x $(BINARY_NAME)

install: build
	@echo "[\033[38;5;82m+\033[0m] Installing to /usr/local/bin..."
	@sudo cp $(BINARY_NAME) /usr/local/bin/

uninstall:
	@echo "[\033[31m!\033[0m] Removing Cyph3r from system..."
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@rm -f $(BINARY_NAME)

clean:
	@go clean
	@rm -f $(BINARY_NAME)
	@rm -rf backups/
