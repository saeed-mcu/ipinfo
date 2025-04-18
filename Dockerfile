# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY main.go .
RUN go build -o ipinfo main.go

FROM alpine:latest

COPY --from=builder /app/ipinfo /ipinfo

EXPOSE 80
ENTRYPOINT ["/ipinfo"]
