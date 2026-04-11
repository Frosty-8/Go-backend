# -------- Build Stage --------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# IMPORTANT: static binary (fixes Alpine issues)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# -------- Runtime Stage --------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]