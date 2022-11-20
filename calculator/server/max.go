package main

import (
	"io"
	"log"

	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	var max int64 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client's stream: %v\n", err)
		}

		if max <= req.Number {
			max = req.Number
		}

		err = stream.Send(&pb.MaxResponse{
			Result: max,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
