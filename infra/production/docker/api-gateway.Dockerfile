FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
WORKDIR /app/services/api-gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/services/api-gateway/api-gateway .
CMD ["./api-gateway"]