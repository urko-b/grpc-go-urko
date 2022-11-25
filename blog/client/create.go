package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/urko-b/grpc-go-urko/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	fmt.Println("---createBlog was invoke---")
	bloge := &pb.Blog{
		AuthorId: "Urko",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), bloge)
	if err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	fmt.Printf("New blog has been created: %s\n", res.Id)
	return res.Id
}
