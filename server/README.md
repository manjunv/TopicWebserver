High Performance server Server written in [Go.](http://golang.org)

## Quickstart

If you just want to start using server, and you have [installed Go](https://golang.org/doc/install) 1.5+ and set your $GOPATH:

Install and run the server server:

```
#go get github.com/topicserver/server
server
```

Install the Curl clients for testing Topicserver

```

## TODO

Add syncronisation routines so that multiple post routines can handle same topic without overwrighting


### Demo

You can connect to a public server server that is running at our demo site: [nats://<hostserverIp>:8080]

### Build

You can build the latest version of the server from the `master` branch. The master branch generally should build and pass tests, but may not work correctly in your environment. Note that stable branches of operating system packagers provided by your OS vendor may not be sufficient.

You need [*Go*](http://golang.org/) version 1.5+ [installed](https://golang.org/doc/install) to build the server server. 

- Run `go version` to verify that you are running Go 1.5+. (Run `go help` for more guidance.)
- Clone the <https://github.com/topicserver/server> repository.
- Run `go build` inside the `/topicserver/server` directory. A successful build produces no messages and creates the server executable `server` in the directory.
- Run `go test ./...` to run the unit regression tests.

## Running

To start the server server with default settings (and no authentication or clustering), you can invoke the `server` binary with no [command line options](#command-line-arguments) or [configuration file](#configuration-file).

```sh
> ./server
starting Webserver for Topics/Blogs handling and POST/GET/DELETE.
```

The server is started and listening for client connections on port 4222 (the default) from all available interfaces. The logs are displayed to stdout as shown above in the server output.

### Clients

Curl is currently supportted clients

### Protocol

The server server uses a [text based protocol] HTTPS/HTTP and confirms to HTTP/1.1 [RFC2616]

### Process Signaling

On Unix systems, the server server responds to the following signals:
| SIGKILL | SIGINT | SIGUSR1 | SIGHUP

The `server` binary can be used to send these signals to running server servers using the `-sl` flag:
