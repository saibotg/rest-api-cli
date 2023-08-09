BINARY_NAME="rest-api-cli"
BIN_DIR="bin"
DIST_DIR="dist"
VERSION="0.0.2"

.PHONY: clean test build package

all: clean test build pack checksum

test:
	$(info **** test ****)
	go test ./...

build:
	$(info **** build ****)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_darwin-amd64/$(BINARY_NAME) main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_linux-amd64/$(BINARY_NAME) main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_windows-amd64/$(BINARY_NAME).exe main.go

pack:
	$(info **** create packages ****)
	tar -czvf $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_darwin-amd64.tar.gz -C $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_darwin-amd64 .
	tar -czvf $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_linux-amd64.tar.gz -C $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_linux-amd64 .
	tar -czvf $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_windows-amd64.tar.gz -C $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_windows-amd64 .

checksum:
	$(info **** create checksums ****)
	sha256sum $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_darwin-amd64.tar.gz > $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_darwin-amd64.tar.gz.sha256.txt
	sha256sum $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_linux-amd64.tar.gz > $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_linux-amd64.tar.gz.sha256.txt
	sha256sum $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_windows-amd64.tar.gz > $(DIST_DIR)/$(BINARY_NAME)_$(VERSION)_windows-amd64.tar.gz.sha256.txt

clean:
	$(info **** clean ****)
	go clean
	rm -rf $(BIN_DIR)/*
	rm -rf $(DIST_DIR)/*
