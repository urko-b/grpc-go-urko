//go:build !test
// +build !test

package main

import (
	"log"
	"net"
	"time"

	pb "github.com/urko-b/grpc-go-urko/greet/proto"

	"google.golang.org/grpc"
)

var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
