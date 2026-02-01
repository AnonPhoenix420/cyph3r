# Step 1: Build the binary
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o cyph3r ./cmd/cyph3r

# Step 2: Create a tiny execution image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cyph3r .
ENTRYPOINT ["./cyph3r"]
