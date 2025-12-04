# Application Configuration Variables
APP_NAME=server
DB_USERNAME=root
DB_PASSWORD=root1234
DB_NAME=scholar_ai
DB_PORT=3306

# Development Commands
# Run the application in development mode with debug logging enabled
# Usage: make dev
dev:
	export GIN_MODE=debug && go run ./cmd/$(APP_NAME)

# Build the application binary
# Compiles the Go application and outputs to bin/$(APP_NAME)
# Usage: make build
build:
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

# Run all tests in verbose mode
# Executes all test files in the project with detailed output
# Usage: make test
test:
	go test -v ./...

# Database Migration Commands (Atlas)
# All migration commands use Atlas with GORM environment configuration
# Migration files are auto-generated with timestamped names (format: YYYYMMDDHHMMSS)

# Generate a new migration file based on schema changes
# Compares current GORM models with database schema and creates migration diff
# Usage: make migrate
migrate:
	atlas migrate diff --env gorm

# Apply all pending migrations to the database
# Updates the database schema to match the latest migration files
# Usage: make up
up:
	atlas schema apply --env gorm -u "mysql://${DB_USERNAME}:${DB_PASSWORD}@localhost:3306/${DB_NAME}"

# Rollback the last migration
# Reverts the most recent migration that was applied to the database
# Usage: make down
down:
	atlas migrate down --env gorm -u "mysql://${DB_USERNAME}:${DB_PASSWORD}@localhost:3306/${DB_NAME}"

# Clean the database schema (WARNING: Destructive operation)
# Removes all tables and data from the database
# Usage: make clean
clean:
	atlas schema clean --env gorm -u "mysql://${DB_USERNAME}:${DB_PASSWORD}@localhost:3306/${DB_NAME}"

# Generate Swagger API documentation
# Scans the codebase for Swagger annotations and generates docs in the docs/ directory
# Usage: make swagger
swagger:
	swag init -g cmd/server/main.go -o docs

# Generate RSA private key and X.509 certificate for JWT signing
# Creates a 2048-bit RSA private key and self-signed certificate in keys/ directory
# Usage: make generate-key
# Options: make generate-key KEY=path/to/key.pem CERT=path/to/cert.pem
generate-key:
	@go run ./cmd/keygen -key $(or $(KEY),keys/private_key.pem) -cert $(or $(CERT),keys/certificate.pem)
	
.PHONY: dev build test migrate up down swagger generate-key