# go-test-servers

Command line utility that implements multiple types of servers for testing comms clients

## Installation
check out this repo and run `go build` in the root directory

## Config
`go-test-servers` uses a yaml config to define which types of servers will be running.  Just define the server parameters as part of the `servers` list

### Example Config
```
servers:
  - type: "socks5"
    enabled: true
    host: 127.0.0.1
    port: 5000
    username: "user"
    password: "password"
    protocol: "tcp"

  - type: "socket"
    enabled: true
    host: 127.0.0.1
    port: 5001

  - type: "ssl-socket"
    enabled: true
    host: 127.0.0.1
    port: 5002
    certfile: "cert.pem"
    keyfile: "key.pem"
    cafile: "ca.pem"
```

## Server Types
This utility currently supports 3 server types

### socks5
a simple socks5 proxy server, which supports username/password authentication.  Currently this proxy only supports the tcp protocol

### socket
a simple socket server that will listen on a port and echo back any data receieved

### ssl-socket
a simple ssl socket server that will listen on a port and echo back any data receieved. `certfile` and `keyfile` are required parameters.  `cafile` is optional.  Server will use system CA certificates in addition to any custom CA provided

## Usage

### Help Statement
```
$ ./go-test-servers --help
Usage of ./go-test-servers:
  -config string
        path to the config file (default "config.yaml")
```

### Example using default config
```
./go-test-servers --config ./config.yaml

2024/08/19 14:07:34 Attempting to start socks5 server
2024/08/19 14:07:34 Enabling authentication
2024/08/19 14:07:34   username: user
2024/08/19 14:07:34   password: password
2024/08/19 14:07:34 socks5 Listening on tcp://127.0.0.1:5000
2024/08/19 14:07:34 Started socks5 server


2024/08/19 14:07:34 Attempting to start socket server
2024/08/19 14:07:34 socket Listening on 127.0.0.1:5001
2024/08/19 14:07:34 Started socket server


2024/08/19 14:07:34 Attempting to start ssl-socket server
2024/08/19 14:07:34 Cannot start ssl server, Cert file not found: cert.pem
2024/08/19 14:07:34 Failed to start ssl-socket server

2024/08/19 14:07:34 Waiting, press Ctrl+C to exit
```
