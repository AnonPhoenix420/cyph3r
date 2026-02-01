# CYPH3R v2.6 Makefile
BINARY_NAME=cyph3r
GO_FILES=cmd/cyph3r/main.go

all: build

build:
	@echo "🛰️  Building CYPH3R Binary..."
	go build -o $(BINARY_NAME) $(GO_FILES)
	chmod +x $(BINARY_NAME)

repair:
	@echo "🔧  Initiating Self-Repair..."
	rm -f go.sum
	go clean -modcache
	go mod tidy
	@echo "🛰️  Rebuilding..."
	go build -o $(BINARY_NAME) $(GO_FILES)
	chmod +x $(BINARY_NAME)
	@echo "[✔] Repair Complete."

clean:
	@echo "🧹  Cleaning workspace..."
	rm -f $(BINARY_NAME)
	rm -f go.sum
	go clean -cache
	@echo "[✔] Workspace pristine."

install: build
	@cp $(BINARY_NAME) /usr/local/bin/ 2>/dev/null || echo "Run 'sudo make install' to move to /usr/local/bin"

.PHONY: all build repair clean install
