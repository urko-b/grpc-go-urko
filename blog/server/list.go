package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/urko-b/grpc-go-urko/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	fmt.Printf("ListBlogs was invoke with %v\n", in)

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %s\n", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		if err = cur.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("cur.Decode Error: %s\n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("cur.Err Error: %s\n", err),
		)
	}

	return nil
}
