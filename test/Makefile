.PHONY: test-unit test-integration test-all

test-unit:
	@echo "Running unit tests..."
	@go test ./...

test-integration:
	@echo "Running integration tests..."
	@go test -tags=integration ./...

test-all:
	@echo "Running all tests..."
	@go test -tags=unit,integration ./...