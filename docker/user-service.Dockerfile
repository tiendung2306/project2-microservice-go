FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/user-service

FROM gcr.io/distroless/static:nonroot AS production

WORKDIR /app

COPY --from=builder /app/user-service /user-service
COPY --from=builder /app/.env /app/.env

EXPOSE 3000

CMD ["/user-service"] 