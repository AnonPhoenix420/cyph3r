# =============================================
# Build Stage
# =============================================
FROM golang:1.23-alpine AS builder

# Install git for dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o cyph3r ./cmd/cyph3r

# =============================================
# Final Minimal Image
# =============================================
FROM alpine:latest

# Install certificates for HTTPS requests
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user for security
RUN adduser -D -u 1000 cyph3ruser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/cyph3r .

# Set permissions
RUN chmod +x ./cyph3r && chown -R cyph3ruser:cyph3ruser /app

# Switch to non-root user
USER cyph3ruser

# Metadata labels
LABEL org.opencontainers.image.title="CYPH3R" \
      org.opencontainers.image.description="Network Intelligence & OSINT Tool" \
      org.opencontainers.image.version="2.6" \
      org.opencontainers.image.authors="AnonPhoenix420"

# Healthcheck (optional but recommended)
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s \
  CMD ./cyph3r --target 8.8.8.8 --full || exit 1

# Default command
ENTRYPOINT ["./cyph3r"]
