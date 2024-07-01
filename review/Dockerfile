# syntax=docker/dockerfile:1

# build stage
FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


COPY ./cmd/coupon_service/main.go ./cmd/coupon_service/main.go
COPY ./internal/ ./internal/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/coupon_service/main.go

#build server
FROM scratch

WORKDIR /

COPY --from=builder /app/server ./

EXPOSE 8080
ENTRYPOINT  ["./server"]
