check:
	@echo "# Checking for suspicious, abnormal, or useless code..."
	@go vet ./...

install:
	@echo "# Installing dependencies..."
	@go mod tidy

tests:
	@echo "# Running tests..."
	@go test -cover ./...

tidy:
	@echo "# Formatting code..."
	@go fmt ./...