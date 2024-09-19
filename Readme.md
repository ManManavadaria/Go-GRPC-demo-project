# Go-GRPC-Demo-Project

## Overview

This project demonstrates a simple implementation of a **gRPC** service using Go. It showcases various types of gRPC communication, such as unary, server-streaming, client-streaming, and bidirectional-streaming RPCs. The project makes use of Protocol Buffers (protobuf) for defining the gRPC service and messages.

## Features

The project implements the following gRPC calls in the `GreetService`:

1. **Unary RPC** - `SayHello`: Sends a single request and receives a single response.
2. **Server-Streaming RPC** - `SayHelloServerStreaming`: Sends a single request and receives multiple responses in a stream.
3. **Client-Streaming RPC** - `SayHelloClientStreaming`: Sends multiple requests in a stream and receives a single response.
4. **Bidirectional-Streaming RPC** - `SayHelloBidirectionalStreaming`: Both client and server send a sequence of messages using a stream.

## Proto File

The **`proto/greet.proto`** file defines the gRPC service and message types used in the project.

### gRPC Service Definition

```proto
syntax = "proto3";

option go_package = "github.com/ManManavadaria/Go-GRPC-demo-project/proto";

package grpc;

service GreetService {
    rpc SayHello(NoParam) returns (HelloResponse);
    rpc SayHelloServerStreaming(NameList) returns (stream HelloResponse);
    rpc SayHelloClientStreaming(stream HelloRequest) returns (MessageList);
    rpc SayHelloBidirectionalStreaming(stream HelloRequest) returns (stream HelloResponse);
}
```

### Message Types

- **NoParam**: Empty request parameter.
- **HelloRequest**: Contains a `message` string sent by the client.
- **HelloResponse**: Contains a `message` string returned by the server.
- **NameList**: Contains a list of `names` (repeated string).
- **MessageList**: Contains a list of `messages` (repeated string).

## Installation

### Prerequisites

- Go (Golang) version 1.16 or later
- Protocol Buffers compiler (protoc) with Go support

### Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/ManManavadaria/Go-GRPC-demo-project.git
   cd Go-GRPC-demo-project
   ```

2. **Install dependencies**:

   Use `go mod` to install the necessary Go modules:

   ```bash
   go mod tidy
   ```

3. **Generate Go code from the proto file**:

   You need to generate Go code from the proto file using the `protoc` compiler.

   Run the following command to generate the Go files:

   ```bash
   protoc --go_out=. --go-grpc_out=. proto/greet.proto
   ```

   This will generate the gRPC service and message structures in the `proto` folder.

## Running the Project

### Server

The server implements the methods defined in the `GreetService`. To run the server:

1. Build the server code:

   ```bash
   go build -o server ./server
   ```

2. Start the server:

   ```bash
   ./server
   ```

   The server will start listening for gRPC calls on a specified port (e.g., `localhost:50051`).

### Client

The client can invoke the different RPC methods (Unary, Server-Streaming, Client-Streaming, Bidirectional-Streaming) from the server. To run the client:

1. Build the client code:

   ```bash
   go build -o client ./client
   ```

2. Run the client:

   ```bash
   ./client
   ```

   You will see the client making gRPC calls to the server and printing the responses.

## Project Structure

```bash
.
├── client/                   # Client-side code for calling the gRPC service
├── proto/                    # Proto file and generated Go code
├── server/                   # Server-side implementation of the gRPC service
├── go.mod                    # Go module definition
├── go.sum                    # Dependencies
└── README.md                 # Project documentation
```

## gRPC Methods Implementation

### 1. Unary RPC - `SayHello`

- **Client** sends an empty request (`NoParam`).
- **Server** responds with a greeting (`HelloResponse`).

### 2. Server-Streaming RPC - `SayHelloServerStreaming`

- **Client** sends a request containing a list of names (`NameList`).
- **Server** responds by sending a greeting message for each name in the stream.

### 3. Client-Streaming RPC - `SayHelloClientStreaming`

- **Client** sends a stream of messages (`HelloRequest`).
- **Server** collects all the messages and responds with a single message list (`MessageList`).

### 4. Bidirectional-Streaming RPC - `SayHelloBidirectionalStreaming`

- **Client** sends a stream of messages (`HelloRequest`).
- **Server** responds with a stream of greetings (`HelloResponse`) for each request.

## How to Customize

- Modify the proto file (`proto/greet.proto`) to add more gRPC methods or messages as needed.
- Re-generate the Go files using `protoc` whenever the proto file is modified.
- Implement additional logic in the server and client code as per your requirements.

## Useful Commands

- **Generate Go files from proto**:
  ```bash
  protoc --go_out=. --go-grpc_out=. proto/greet.proto
  ```
  
- **Build server**:
  ```bash
  go build -o server ./server
  ```

- **Build client**:
  ```bash
  go build -o client ./client
  ```

---
