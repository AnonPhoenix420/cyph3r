# Build Stage
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download && go build -o cyph3r ./cmd/cyph3r

# Final Tactical Image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/cyph3r .
RUN chmod +x ./cyph3r

# Execution
ENTRYPOINT ["././cyph3r"]
