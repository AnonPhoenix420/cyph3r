# STAGE 1: Build the binary
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git and certs (needed for Go modules and HTTPS API calls)
RUN apk add --no-cache git ca-certificates

# Copy go.mod and go.sum (if they exist) and download dependencies
COPY go.mod* go.sum* ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cyph3r ./cmd/cyph3r/main.go

# STAGE 2: Run the binary
FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/cyph3r .

# Command to run (allows passing flags like -target)
ENTRYPOINT ["./cyph3r"]
