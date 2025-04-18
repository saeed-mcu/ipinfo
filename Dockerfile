# Build stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY main.go .
RUN go build -o ipinfo main.go

# Final image (very small)
FROM alpine:latest

COPY --from=builder /app/ipinfo /ipinfo

EXPOSE 80
ENTRYPOINT ["/ipinfo"]
