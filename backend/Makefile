.PHONY: help build run test migrate-up migrate-down migrate-create clean docker-up docker-down

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o bin/server ./cmd/server
	go build -o bin/migrate ./cmd/migrate

run: ## Run the application
	go run ./cmd/server

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

migrate: ## Run database migrations (code-first: uses GORM AutoMigrate)
	go run ./cmd/migrate

seed: ## Seed database with test data
	go run ./cmd/seed

seed-reset: ## Reset and reseed database
	go run ./cmd/seed -reset

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html

docker-up: ## Start Docker containers
	docker-compose up -d

docker-down: ## Stop Docker containers
	docker-compose down

docker-logs: ## View Docker logs
	docker-compose logs -f

install-deps: ## Install dependencies
	go mod download
	go mod tidy

lint: ## Run linter
	golangci-lint run

