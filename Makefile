APP_NAME=myapp
SRC=main.go
OUTPUT_DIR=bin

# Default target
all: build-linux build-mac build-mac-arm build-windows

# Linux (amd64)
build-linux:
	@echo "🔧 Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$(APP_NAME)-linux $(SRC)

# macOS (Intel)
build-mac:
	@echo "🔧 Building for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$(APP_NAME)-mac $(SRC)

# macOS (Apple Silicon)
build-mac-arm:
	@echo "🔧 Building for macOS (arm64)..."
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$(APP_NAME)-mac-arm64 $(SRC)

# Windows (64-bit)
build-windows:
	@echo "🔧 Building for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$(APP_NAME)-win.exe $(SRC)

# Clean binaries
clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(OUTPUT_DIR)

# Create output dir if not exist
$(shell mkdir -p $(OUTPUT_DIR))