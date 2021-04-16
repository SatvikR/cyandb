# CyanDB

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SatvikR/cyandb)](https://pkg.go.dev/github.com/SatvikR/cyandb)

## About

This is a lightweight key-value database written in go. CyanDB is persistent,
and is not an in-memory database.

## Project Status

This project is in early stages of development and should not be used in production

## To run

- Server:

```bash
# windows
cyand 
# On linux/mac
sudo cyand
```

- Client

```bash
# windows/linux/mac
cyansh
```

## To install:

For all platforms: Make sure that your `GOBIN` directory, typically `~/go/bin/` is in your `PATH`

- Windows:

```bash
go get github.com/SatvikR/cyandb/cmd/cyand
go get github.com/SatvikR/cyandb/cmd/cyansh
```

- Linux/Mac:

```bash
go get github.com/SatvikR/cyandb/cmd/cyand
go get github.com/SatvikR/cyandb/cmd/cyansh

sudo cp $(which cyand) /usr/bin
```

## To compile from source:

From the root of the repository, run the following:

- Windows:
  `.\make.bat install`
  
- Linux:
  `make install`


## To compile for development:

From the root of the repository, run the following:

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

## LICENSE

[Apache 2.0](https://github.com/SatvikR/cyandb/blob/master/LICENSE)

## FAQ

- Where does the name CyanDB come from?
  - I was originally going to call this project BlueDB, but apparently that already existed.
    When I asked one of my friends, [Alexandre](https://github.com/Alexandre2006), "What should I rename BlueDB to?", he replied "CyanDB."

