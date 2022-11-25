package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/urko-b/grpc-go-urko/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("doGreetWithDeadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Urko",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Printf("Unexpected gRPC error: %s\n", err)
			}
		} else {
			log.Fatalf("non gRPC error: %s\n", err)
		}
		log.Fatalf("%s", err)
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
