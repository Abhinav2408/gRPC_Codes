package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{N: 120}

	stream, err := c.Prime(context.Background(), req)

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

		log.Printf("Message received: %v\n", msg.Res)
	}
}
