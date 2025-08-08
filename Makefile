# Makefile for Go project with DB operations
BINARY_NAME=myapp
MIGRATION_DIR=db/migrations
SEED_DIR=db/seeds

# Load .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Build application
build:
	@go build -o $(BINARY_NAME) .

log:
	@echo $(DB_URL)

# Run application
run: build
	@./$(BINARY_NAME)

# Clean build
clean:
	@rm -f $(BINARY_NAME)

# Database migrations
db-migrate:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DBSTRING) goose -dir $(MIGRATION_DIR) up

db-migrate-rollback:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DBSTRING) goose -dir $(MIGRATION_DIR) down

# Database seeding
db-seed:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DBSTRING)" goose -dir $(SEED_DIR) up

db-seed-rollback:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DBSTRING)" goose -dir $(SEED_DIR) down-to 0

# Show migration status
db-status:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DBSTRING)" goose -dir $(MIGRATION_DIR) status

# Help message
help:
	@echo "Available commands:"
	@echo "  make build               - Build application"
	@echo "  make run                 - Run application"
	@echo "  make clean               - Clean build"
	@echo ""
	@echo "Database operations:"
	@echo "  make db-migrate          - Run all pending migrations"
	@echo "  make db-migrate-rollback - Rollback last migration"
	@echo "  make db-seed             - Run all seed files"
	@echo "  make db-seed-rollback    - Rollback all seeds"
	@echo "  make db-status           - Show migration status"
