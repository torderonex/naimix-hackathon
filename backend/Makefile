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