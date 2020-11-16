# CyanDB

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SatvikR/cyandb)](https://pkg.go.dev/github.com/SatvikR/cyandb@v0.1.0)

## About 

This is a simple key-value database written in go. CyanDB is persistent,
and is not an in-memory database.

## Project Status

This project is in early stages of development and should not be used in production

## TODO

- [x] Rewrite serialization
- [x] Rewrite Get command
- [ ] Make Set command rewrite existing key
- [ ] Add error handling
- [ ] Add websocket functionality to server
- [ ] Create the client
- [ ] Create CLI to start program, eg. `cyan run server` and `cyan run shell`

## LICENSE

[Apache 2.0](https://github.com/SatvikR/cyandb/blob/master/LICENSE)

## FAQ

- Where does the name CyanDB come from?
    - I was originally going to call this project BlueDB, but apparently that already existed. 
    When I asked one of my friends, [Alexandre2006](https://github.com/Alexandre2006), "What should I rename BlueDB to?", he replied "CyanDB."
