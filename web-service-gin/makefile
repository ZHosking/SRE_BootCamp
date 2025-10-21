# Makefile for Student REST API

# Go variables
APP_NAME = student-api
MAIN_FILE = main.go

# Default target (if you just run `make`)
default: build

# Build the binary
build:
	go build -o $(APP_NAME) $(MAIN_FILE)

# Run the application (for WSL/Linux/macOS)
run:
	@echo "Starting the Student API (Unix/WSL)..."
	@if [ -f .env ]; then export $$(grep -v '^#' .env | xargs); fi; \
	go run $(MAIN_FILE)

# Run the application (for Windows PowerShell)
run-win:
	@echo "Starting the Student API (Windows PowerShell)..."
	powershell -Command "if (Test-Path .env) { Get-Content .env | ForEach-Object { if ($_ -notmatch '^#' -and $_) { $parts = $_ -split '='; [System.Environment]::SetEnvironmentVariable($parts[0], $parts[1]) } } }; go run $(MAIN_FILE)"

# Run tests
test:
	go test ./... -v

# Format code
fmt:
	go fmt ./...

# Clean up built binary
clean:
	rm -f $(APP_NAME)

# Run DB migrations (optional)
migrate:
	go run scripts/migrate.go
