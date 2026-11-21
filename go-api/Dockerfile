# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /build

RUN go mod init go-hello-api

COPY main.go .

RUN go build -o app .

# Runtime stage
FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 8080

USER appuser

CMD ["./app"]