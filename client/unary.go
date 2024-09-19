package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ManManavadaria/Go-GRPC-demo-project/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", res.Message)
}
