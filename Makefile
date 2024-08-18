PROJECT_NAME=ascii-art-gen

build:
		@echo "Building the project..."
		@go build -o ./$(PROJECT_NAME) ./main.go
		@echo "Build successful!"

fmt:
		@echo "Formatting the code..."
		@go fmt ./...

test:
		@echo "Running tests..."
		@go test -v ./...

release-check:
		goreleaser check

release-snapshot:
		goreleaser release --snapshot --clean
