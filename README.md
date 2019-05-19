# Payments

gRPC client and server demo built in GO.

It aims to mimic the behaviour of a payment API and its client.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites to develop

**Install GO** (https://golang.org/doc/install)

**Install gRPC**

```
go get -u google.golang.org/grpc
```

**Install Protocol Buffers v3**

Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(protoc-<version>-<platform>.zip) from here: https://github.com/google/protobuf/releases

Unzip this file.
Update the environment variable PATH to include the path to the protoc binary file.

**Install the protoc plugin for Go**

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.

```
export PATH=$PATH:$GOPATH/bin
```

### Try it

Clone this repository

```
https://github.com/afasola/payments.git
```

run the server (default port is 20000)

```
go run server/server.go
```

### Docker

Build the container

```
docker build -t afasola/payments .
```

Run the container 

```
docker run -p 20000:20000 afasola/payments
```

## Authors

* **Andrea Fasola (ciccio)** 

