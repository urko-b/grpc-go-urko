package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	fmt.Println("doPrimes was invoked")

	stream, err := c.Prime(context.Background(), &pb.PrimeRequest{
		Number: 20,
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	factors := []string{}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		factors = append(factors, fmt.Sprint(msg.Result))
	}
	fmt.Printf("(%s)\n", strings.Join(factors, ","))
}
