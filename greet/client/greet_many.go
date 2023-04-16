package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/greet/proto"
)

func doGreetMany(c pb.GreetServiceClient) {
	log.Println("doGreetMany was invoked")

	req := &pb.GreetRequest{FirstName: "Abhinav"}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error found")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error occured")
		}

		log.Printf("Message received: %v\n", msg.Result)
	}
}
