.PHONY: client

client:
	@echo "Building the client binary"
	go build -o bin/client cmd/client/main.go