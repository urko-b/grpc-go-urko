package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/urko-b/grpc-go-urko/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("---deleteBlog was invoke---")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("c.DeleteBlog: %s", err)
	}

	fmt.Println("Blog was delete")
}
