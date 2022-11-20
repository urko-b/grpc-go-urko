package main

import (
	"io"
	"log"

	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	var sum int64 = 0
	var count int64 = 1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		count++
		sum += req.Number
	}
}
