package main

import (
	"context"
	"io"
	"log"

	pb "github.com/urko-b/grpc-go-urko/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Urko",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
