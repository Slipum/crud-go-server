FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod ./
# Можно раскомментировать строку ниже, если появится go.sum
# COPY go.sum ./
RUN go mod download
COPY . .
# запуск тестов перед компиляцией
RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app .



FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
