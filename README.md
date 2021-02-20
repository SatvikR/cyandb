# CyanDB

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SatvikR/cyandb)](https://pkg.go.dev/github.com/SatvikR/cyandb)

## About

This is a simple key-value database written in go. CyanDB is persistent,
and is not an in-memory database.

## Project Status

This project is in early stages of development and should not be used in production

## To install:

- All platforms: `go get github.com/SatvikR/cyandb`

> To run the server on linux/mac, you need to run with sudo perms:
> `sudo cyandb start server`

## Linux Issues:

Since the server requires root access, if your `GOBIN` directory, typically `/home/yourname/go/bin` is not a part of the
`secure_path` option in `/etc/sudoers`, the command: `sudo cyandb start server` will return an error. To work around this, you can run this:

```bash
sudo cp $(which cyandb) /usr/bin
```



## To compile for development:

- Windows:
  `.\make.bat`
- Linux: 
    `make`

Binaries will show up in `./bin` folder 

## To clean binaries:
- Windows:
  `.\make.bat clean`
- Linux:
  `make clean`

## TODO

- [x] Rewrite serialization
- [x] Rewrite Get command
- [x] Make Set command rewrite existing key
- [x] Fix Get command not finding last key if the key is the last key in the file.
- [x] Add error handling
- [x] Add websocket functionality to server
- [x] Create the client
- [x] Create CLI to start program, e.g. `cyan run server` and `cyan run shell`
- [x] Create command parser on server
- [x] Make client wait for response before reprinting shell

## LICENSE

[Apache 2.0](https://github.com/SatvikR/cyandb/blob/master/LICENSE)

## FAQ

- Where does the name CyanDB come from?
  - I was originally going to call this project BlueDB, but apparently that already existed.
    When I asked one of my friends, [Alexandre2006](https://github.com/Alexandre2006), "What should I rename BlueDB to?", he replied "CyanDB."

