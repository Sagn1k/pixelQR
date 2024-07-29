# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=pixelQR
BINARY_LINUX_AMD64=$(BINARY_NAME)_linux_amd64
BINARY_LINUX_ARM64=$(BINARY_NAME)_linux_arm64
BINARY_MAC_AMD64=$(BINARY_NAME)_mac_amd64
BINARY_MAC_ARM64=$(BINARY_NAME)_mac_arm64

# Directories
SRC_DIR=./cmd/server

# Default target executed when you type 'make'
all: test build-linux-amd64 build-linux-arm64 build-mac-amd64 build-mac-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX_AMD64) $(SRC_DIR)/main.go
	@echo "Built Linux AMD64 binary: $(BINARY_LINUX_AMD64)"

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BINARY_LINUX_ARM64) $(SRC_DIR)/main.go
	@echo "Built Linux ARM64 binary: $(BINARY_LINUX_ARM64)"

build-mac-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_MAC_AMD64) $(SRC_DIR)/main.go
	@echo "Built macOS AMD64 binary: $(BINARY_MAC_AMD64)"

build-mac-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BINARY_MAC_ARM64) $(SRC_DIR)/main.go
	@echo "Built macOS ARM64 binary: $(BINARY_MAC_ARM64)"

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_LINUX_AMD64)
	rm -f $(BINARY_LINUX_ARM64)
	rm -f $(BINARY_MAC_AMD64)
	rm -f $(BINARY_MAC_ARM64)

run: build-linux-amd64
	./$(BINARY_LINUX_AMD64)

run-mac: build-mac-arm64
	./$(BINARY_MAC_ARM64)

docker-build:
	docker build -t pixelqr:latest .

docker-run:
	docker run -p 3000:3000 pixelqr:latest

.PHONY: all build-linux-amd64 build-linux-arm64 build-mac-amd64 build-mac-arm64 clean test run docker-build docker-run
