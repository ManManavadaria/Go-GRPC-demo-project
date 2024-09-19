package main

import (
	"log"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Jhon", "Alice", "Bob"},
	}

	callSayHello(client)

	callSayHelloServerStreaming(client, names)

	callSayHelloClientStreaming(client, names)

	callHelloBidirectionalStreaming(client, names)
}
