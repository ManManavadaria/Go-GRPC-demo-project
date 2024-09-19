package main

import (
	"io"
	"log"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(clientStream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req, err := clientStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name : %v", req.Message)

		res := &pb.HelloResponse{
			Message: "Hello " + req.Message,
		}

		if err := clientStream.Send(res); err != nil {
			return err
		}
	}
}
