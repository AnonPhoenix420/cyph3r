build:
	go mod tidy
	go build -o cyph3r ./cmd/cyph3r/main.go

run:
	./cyph3r -target google.com

clean:
	rm -f cyph3r
