GO=go
BIN=./bin
ARGS=-i

# Only to be used in development
build:
	$(GO) build $(ARGS) -o $(BIN)/cyan cyan.go

clean:
	rm -rf $(BIN)