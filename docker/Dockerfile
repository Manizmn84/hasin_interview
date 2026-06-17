FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o hasin-backend ./cmd/app/main.go

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app


COPY --from=builder /app/hasin-backend .
COPY --from=builder /app/.env .

COPY internal/infrastructure/jwt/ ./internal/infrastructure/jwt/

EXPOSE 8080


CMD ["./hasin-backend"]