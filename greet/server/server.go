package main

import pb "github.com/urko-b/grpc-go-urko/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
