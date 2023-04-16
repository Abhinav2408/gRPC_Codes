package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlog(ctx context.Context, stream pb.BlogService_ListBlogsServer) error {
	log.Println("Listblog invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			"internal error",
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}

		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				"internal error",
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			"internal error",
		)
	}

	return nil
}
