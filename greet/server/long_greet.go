package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/urko-b/grpc-go-urko/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet invoked with")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error: %v \n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
