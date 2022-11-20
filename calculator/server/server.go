package main

import pb "github.com/urko-b/grpc-go-urko/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
