FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/auth-service

FROM gcr.io/distroless/static:nonroot AS production

WORKDIR /app

COPY --from=builder /app/auth-service /auth-service

COPY --from=builder /app/.env /app/.env

EXPOSE 3001

CMD ["/auth-service"]