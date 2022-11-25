package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/urko-b/grpc-go-urko/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
	fmt.Printf("DeleteBlog was invoke with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("primitive.ObjectIDFromHex: %s", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("collection.DeleteOne: %s", err),
		)
	}
	if res.DeletedCount <= 0 {
		return &emptypb.Empty{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("blog not found with ID: %s", in.Id),
		)
	}
	fmt.Printf("blog was deleted: %s\n", in.Id)

	return &emptypb.Empty{}, nil
}
