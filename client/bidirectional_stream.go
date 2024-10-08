package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"
)

func callHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Print("Starting bidirectional streaming")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send name : %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming : %v", err)
			}
			log.Print(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc

	log.Print("Bidirectional streaming finished")
}
