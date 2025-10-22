# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=dhis2
BINARY_UNIX=$(BINARY_NAME)_unix
INSTALL_DIR=$(HOME)/bin
COMPLETION_FILE=$(HOME)/.$(BINARY_NAME)_completion.zsh
SHELL=zsh
SHELLRC_FILE=$(HOME)/.$(SHELL)rc
# BUILD_VERSION=$(./get-version.sh)
CURRENT_DIR=$(pwd)
VERSION := $(shell ./get-version.sh)
SOURCE_DIR := .
#ZSHRC_FILE=$(HOME)/.zshrc

# All target: compile and build
all: test build

# Build the project
build:
	echo "version is $(VERSION)"
	$(GOBUILD) -ldflags "-X cmd.version=$(VERSION)" -o $(BINARY_NAME) $(SOURCE_DIR)

# Run the tests
test:
	$(GOTEST) -v ./...

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run the project
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) -v ./...

# Cross compile for Unix
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-s -w -X cmd.version=$(VERSION)" -o $(BINARY_UNIX) -v

# Cross compile for Unix
build-macos:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags "-s -w -X cmd.version=$(VERSION)" -o $(BINARY_UNIX) -v

# Cross compile for Windows
build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags "-s -w -X cmd.version=$(VERSION)" -o $(BINARY_NAME).exe -v

# Install the binary to $HOME/bin
install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)
	echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)"
	# $(INSTALL_DIR)/$(BINARY_NAME) completion zsh >> $(COMPLETION_FILE)
	echo "Generated completion script for $(BINARY_NAME) and appended to $(SHELLRC_FILE)"
	grep -qxF 'source <($(INSTALL_DIR)/$(BINARY_NAME) completion zsh)' $(SHELLRC_FILE) || echo 'source <($(INSTALL_DIR)/$(BINARY_NAME) completion zsh)' >> $(SHELLRC_FILE)