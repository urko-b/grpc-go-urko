package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/urko-b/grpc-go-urko/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	fmt.Println("---listBlog was invoke---")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("c.ListBlogs: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv: ", err)
		}

		fmt.Printf("%v\n", res)
	}
}
