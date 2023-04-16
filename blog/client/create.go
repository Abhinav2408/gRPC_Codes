package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Abhinav",
		Title:    "First Blog",
		Content:  "First Content",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Blog has been created : %s\n", res.Id)
	return res.Id
}
