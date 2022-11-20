package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	fmt.Println("doAvg was invoked")
	reqs := []*pb.AvgRequest{
		{
			Number: 526,
		},
		{
			Number: 15,
		},
		{
			Number: 5,
		},
		{
			Number: 4345,
		},
		{
			Number: 567,
		},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)

		stream.Send(req)
		time.Sleep(200 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Avg: %f\n", res.Result)
}
