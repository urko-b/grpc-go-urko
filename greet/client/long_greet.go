package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/urko-b/grpc-go-urko/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{
			FirstName: "Urko",
		},
		{
			FirstName: "Yan",
		},
		{
			FirstName: "Lebron",
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("LongGreet: %s\n", res.Result)
}
