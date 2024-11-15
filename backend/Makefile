include ./config/.env

MIGRATION_DIR=db/migrations

DB_DSN=postgresql://$(PG_USR):$(PG_PWD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable


.PHONY: run
run:
	@echo "Running the application..."
	go run ./cmd/app/main.go

.PHODY: build
build:
	@echo "Building the application.."
	go build -o ./bin/ ./cmd/app/main.go
	@echo "Binary successfuly build to ./bin/main"
.PHODY: swag
swag:
	@echo "Generating swagger docs.."
	swag init -d="./cmd/app,./internal/handler,./internal/entities"

.PHONY: migrate-up
migrate-up:
	@echo "Applying migrations..."
	goose -dir $(MIGRATION_DIR) postgres "$(DB_DSN)" up

.PHONY: migrate-down
migrate-down:
	@echo "Reverting the last migration..."
	goose -dir $(MIGRATION_DIR) postgres "$(DB_DSN)" down

.PHONY: migrate-create
migrate-create:
	@echo "Creating new migrate file..."
	goose -dir $(MIGRATION_DIR) postgres "$(DB_DSN)" create $(name) sql