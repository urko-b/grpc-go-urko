package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	reqs := []*pb.MaxRequest{
		{
			Number: 4,
		},
		{
			Number: 41,
		},
		{
			Number: 442,
		},
		{
			Number: 455,
		},
		{
			Number: 4123,
		},
		{
			Number: 422,
		},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("stream.Recv: %v", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
