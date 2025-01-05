.PHONY: all run doc

APP_NAME=vehicle-registration-manager
DOCS_DIR=docs
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://yourusername:yourpassword@localhost:5432/dealership?search_path=prod
GOOSE_MIGRATION_DIR=./migrations
APP_NAME_LOGGER=vehicle-registration-manager
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=dealership
DRIVER_NAME=postgres
DB_SCHEMA=prod
SCOPE=prod
SCOPE_LEVEL_LOGGER=info

all: run

run:
	@echo "Running the application..."
	APP_NAME_LOGGER=$(APP_NAME_LOGGER) \
	DB_HOST=$(DB_HOST) \
	DB_PORT=$(DB_PORT) \
	DB_USER=$(DB_USER) \
	DB_PASSWORD=$(DB_PASSWORD) \
	DB_NAME=$(DB_NAME) \
	DRIVER_NAME=$(DRIVER_NAME) \
	DB_SCHEMA=$(DB_SCHEMA) \
	SCOPE=$(SCOPE) \
	SCOPE_LEVEL_LOGGER=$(SCOPE_LEVEL_LOGGER) \
	go run cmd/$(APP_NAME)/main.go

doc:
	@echo "Generating API documentation..."
	swag init -g cmd/$(APP_NAME)/main.go -o $(DOCS_DIR)

clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(DOCS_DIR)/*

docker-up-all:
	@echo "Starting the application with Docker Compose..."
	docker-compose up --build -d

docker-down-all:
	@echo "Stopping the application with Docker Compose..."
	docker-compose down

docker-up-database:
	@echo "Starting the database with Docker Compose..."
	docker-compose up --build -d postgres

docker-down-database:
	@echo "Stopping the database with Docker Compose..."
	docker-compose down postgres

docker-up-app:
	@echo "Starting the application with Docker Compose..."
	docker-compose up --build -d app

docker-down-app:
	@echo "Stopping the application with Docker Compose..."
	docker-compose down app

goose-up:
	@echo "Running database migrations..."
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	goose up

goose-status:
	@echo "Running database migrations status..."
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	goose status

goose-down:
	@echo "Rolling rollback database migrations..."
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	goose down

help:
	@echo "Makefile commands:"
	@echo "  run   - Run the application"
	@echo "  doc   - Generate API documentation"
	@echo "  clean - Clean up generated files"
	@echo "  docker-up-all    - Start the application with Docker Compose"
	@echo "  docker-down-all  - Stop the application with Docker Compose"
	@echo "  docker-up-database    - Start the database with Docker Compose"
	@echo "  docker-down-database  - Stop the database with Docker Compose"
	@echo "  docker-up-app    - Start the application with Docker Compose"
	@echo "  docker-down-app  - Stop the application with Docker Compose"
	@echo "  goose-status- Check database migrations status"
	@echo "  goose-up    - Run database migrations"
	@echo "  goose-down  - Rollback database migrations"
	@echo "  help  - Show this help message"