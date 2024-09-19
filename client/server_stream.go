package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Print("Server Streaming started")
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*8)
	defer cancle()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(message.Message)
	}
	log.Print("Server Streaming finished")
}
