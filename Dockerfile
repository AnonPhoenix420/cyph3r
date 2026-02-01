# Build Stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install build essentials
RUN apk add --no-cache git make

# Copy source code
COPY . .

# Build the tool
RUN go mod tidy
RUN go build -o cyph3r ./cmd/cyph3r

# Final Stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/cyph3r .

# Command to run
ENTRYPOINT ["./cyph3r"]
