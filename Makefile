GO=go
BIN=./bin
ARGS=-i

# Only to be used in development
build:
	$(GO) build $(ARGS) -o $(BIN)/cyandb cyan.go

clean:
	rm -rf $(BIN)

install:
	$(GO) install
	sudo cp $$(which cyandb) /usr/bin/cyandb
