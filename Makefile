.PHONY: migrate run test

# Migrate database
migrate:
    go run cmd/migrate/main.go --migrate

# Run the application
run:
    go run cmd/api/main.go

# Run tests
test:
    go test ./...