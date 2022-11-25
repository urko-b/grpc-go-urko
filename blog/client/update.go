package main

import (
	"context"
	"fmt"

	pb "github.com/urko-b/grpc-go-urko/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Urko",
		Title:    "A new title",
		Content:  "Content of the first blog with some additions",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		fmt.Printf("c.UpdateBlog: %s\n", err)
	}

	fmt.Println("Blog was updated")
}
