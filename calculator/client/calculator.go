package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient, x int32, y int32) {
	log.Println("doCalculator was invoked")

	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		One: x,
		Two: y,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
