package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("DeleteBlog invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"cannot parse id",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"cannot delete",
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.Internal,
			"no blog with id found",
		)
	}

	return &emptypb.Empty{}, nil
}
