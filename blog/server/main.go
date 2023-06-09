package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Abhinav2408/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "localhost:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = (client.Database("blogdb")).Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
