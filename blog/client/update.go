package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("update blog invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Abhinav",
		Title:    "New Title",
		Content:  "New Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Blog Updated")
}
