# CyanDB

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SatvikR/cyandb)](https://pkg.go.dev/github.com/SatvikR/cyandb)

## What is this

A simple key-value database written in go

## About

This project is in early stages of development and should not be used in production

## TODO

- [x] Rewrite serialization
- [ ] Make Set command rewrite existing key
- [ ] Rewrite Get command
- [ ] Add error handling
- [ ] Add websocket functionality to server
- [ ] Create the client
- [ ] Create CLI to start program, eg. `cyan run server` and `cyan run shell`

## FAQ

- Where does the name CyanDB come from?
    - I was originally going to call this project BlueDB, but apparently that already existed. 
    When I asked one of my friends, [Alexandre2006](https://github.com/Alexandre2006), "What should I rename BlueDB to?", he replied "CyanDB."