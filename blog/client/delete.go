package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Blog deleted")
}
