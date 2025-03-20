FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
WORKDIR /app/services/trip-service
RUN CGO_ENABLED=0 GOOS=linux go build -o trip-service ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/services/trip-service/trip-service .
CMD ["./trip-service"] 