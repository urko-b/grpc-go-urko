package main

import (
	"context"
	"fmt"

	pb "github.com/urko-b/grpc-go-urko/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	fmt.Println("---readBlog was invoked---")

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		fmt.Printf("c.ReadBlog: %s\n", err)
	}

	fmt.Printf("Blog: %v\n", res)
	return res
}
