package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, name *pb.NameList) {
	log.Println("Client streaming started")

	stream, err := client.SayHelloClentStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names : %v", err)
	}

	for _, name := range name.Names {
		req := &pb.HelloRequest{
			Message: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Print("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
