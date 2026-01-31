all: repair

build:
	go build -o cyph3r ./cmd/cyph3r

repair:
	@echo "[!] Cleaning cache and fixing go.sum..."
	go clean -modcache
	rm -f go.sum
	go mod tidy
	go mod verify
	go build -o cyph3r ./cmd/cyph3r
	@echo "[✔] Done. Run with ./cyph3r"
