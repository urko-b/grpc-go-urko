package main

import (
	"context"
	"log"
	"net"

	pb "github.com/urko-b/grpc-go-urko/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogServiceServer
}

var (
	addr       string = "0.0.0.0:50051"
	collection *mongo.Collection
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("mongo.NewClient: %v\n", err)
	}

	if err = client.Connect(context.Background()); err != nil {
		log.Fatalf("client.Connect: %v\n", err)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
