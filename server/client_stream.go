package main

import (
	"io"
	"log"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"
)

func (s *helloServer) SayHelloClentStreaming(stream pb.GreetService_SayHelloClentStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			log.Fatal(err)
			return err
		}

		log.Printf("Got request with name : %v", req.Message)

		messages = append(messages, "Hello "+req.Message)
	}
}
