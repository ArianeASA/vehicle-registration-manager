.PHONY: all run doc

APP_NAME=vehicle-registration-manager
DOCS_DIR=docs

all: run

run:
	@echo "Running the application..."
	go run cmd/$(APP_NAME)/main.go

doc:
	@echo "Generating API documentation..."
	swag init -g cmd/$(APP_NAME)/main.go -o $(DOCS_DIR)

clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(DOCS_DIR)/*

docker-up:
	@echo "Starting the application with Docker Compose..."
	docker-compose up --build -d

docker-down:
	@echo "Stopping the application with Docker Compose..."
	docker-compose down

help:
	@echo "Makefile commands:"
	@echo "  run   - Run the application"
	@echo "  doc   - Generate API documentation"
	@echo "  clean - Clean up generated files"
	@echo "  help  - Show this help message"