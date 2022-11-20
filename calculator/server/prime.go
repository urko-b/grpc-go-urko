package main

import (
	pb "github.com/urko-b/grpc-go-urko/calculator/proto"
)

func (*Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	var k int64 = 2
	n := in.Number
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n = n / k
		} else {
			k++
		}
	}
	return nil
}
