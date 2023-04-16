package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("listBlog was invoked")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(res)
	}
}
