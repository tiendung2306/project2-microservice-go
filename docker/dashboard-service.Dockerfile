FROM golang:1.24.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o dashboard-service ./cmd/dashboard-service

EXPOSE 3004

CMD ["./dashboard-service"] 