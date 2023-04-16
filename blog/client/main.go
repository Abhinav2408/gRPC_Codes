package main

import (
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//second paramter is for security
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)

}
