# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Main package
MAIN_PACKAGE=./cmd/main.go

# Name of the binary
BINARY_NAME=app

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	./$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
