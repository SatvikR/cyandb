GO=go
BIN=./bin
ARGS=-i

build:
    $(GO) build $(ARGS) ./cmd/cyand/main.go -o $(BIN)/cyand
    $(GO) build $(ARGS) ./cmd/cyansh/main.go -o $(BIN)/cyansh

clean:
	rm -rf $(BIN)

install:
	$(GO) install $(ARGS) ./cmd/cyand
	$(GO) install $(ARGS) ./cmd/cyansh
	sudo cp $$(which cyandb) /usr/bin/cyandb
