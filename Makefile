# CYPH3R PROJECT GHOST_BUILD
BINARY_NAME=cyph3r
CMD_PATH=./cmd/cyph3r

.PHONY: all build clean scrub

all: build

build:
	@echo "\033[38;5;39m[*] Compiling Cyph3r Tactical Engine...\033[0m"
	@go build -o $(BINARY_NAME) $(CMD_PATH)
	@chmod +x $(BINARY_NAME)
	@echo "\033[38;5;82m[+] Build Complete: ./$(BINARY_NAME)\033[0m"

clean:
	@rm -f $(BINARY_NAME)
	@go clean

scrub:
	@echo "\033[38;5;196m[!] WARNING: PERFORMING OPSEC SCRUB...\033[0m"
	@rm -f $(BINARY_NAME)
	@history -c && history -w
	@echo "\033[38;5;82m[+] System Scrubbed. No local trace remains.\033[0m"
