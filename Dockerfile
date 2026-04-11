# Stage 1: Build
FROM golang:1.26 AS builder

WORKDIR /app

# Copy modules first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o main .

# Stage 2: Minimal runtime
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run app
CMD ["./main"]