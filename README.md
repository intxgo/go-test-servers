# go-test-servers

Command line utility that implements multiple types of servers for testing comms clients

## Installation
check out this repo and run `go build` in the root directory

## Usage
```
./go-test-servers -h
A command line tool for running test servers

Usage:
  go-test-server [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  socks5      Run a SOCKS5 Server
  ssl         Run a SSL Socket Server (echo server)
  tcp         Run a TCP Socket Server (echo server)

Flags:
  -h, --help   help for go-test-server

Use "go-test-server [command] --help" for more information about a command.
```
