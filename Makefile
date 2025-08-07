PROJECT_NAME=ascii-art

build:
		@echo "Building the project..."
		@go build -o ./dist/$(PROJECT_NAME) ./main.go
		@echo "Build successful!"

fmt:
		@echo "Formatting the code..."
		@go fmt ./...

test:
		@echo "Running tests..."
		@go test -v ./...
