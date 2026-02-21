# CYPH3R PROJECT GHOST_BUILD
BINARY_NAME=cyph3r
CMD_PATH=./cmd/cyph3r
OUT_DIR=bin

.PHONY: all build clean install scrub

all: build

build:
	@echo "\033[38;5;39m[*] Compiling Cyph3r Tactical Engine...\033[0m"
	@mkdir -p $(OUT_DIR)
	@go build -o $(OUT_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "\033[38;5;82m[+] Build Complete: $(OUT_DIR)/$(BINARY_NAME)\033[0m"

install:
	@echo "\033[38;5;39m[*] Installing to /usr/local/bin...\033[0m"
	@sudo cp $(OUT_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "\033[38;5;82m[+] Installation Successful.\033[0m"

clean:
	@echo "\033[38;5;214m[*] Removing temporary build files...\033[0m"
	@rm -rf $(OUT_DIR)
	@go clean

# OPSEC SCRUB: Wipes binary and shell history to prevent forensic trace
scrub:
	@echo "\033[38;5;196m[!] WARNING: PERFORMING OPSEC SCRUB...\033[0m"
	@rm -rf $(OUT_DIR)
	@rm -f /usr/local/bin/$(BINARY_NAME)
	@history -c && history -w
	@echo "\033[38;5;82m[+] System Scrubbed. No local trace remains.\033[0m"
