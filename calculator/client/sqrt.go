package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	fmt.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			fmt.Printf("Error message from server: %s\n", e.Message())
			fmt.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				fmt.Printf("We probably sent a negative number!\n")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}
	fmt.Printf("Sqrt: %f\n", res.Result)
}
