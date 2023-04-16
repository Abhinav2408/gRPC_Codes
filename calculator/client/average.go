package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage invoked")

	reqs := []*pb.AverageRequest{
		{Num: 3},
		{Num: 5},
		{Num: 7},
		{Num: 10},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	for _, req := range reqs {
		stream.Send(req)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Average is %v\n", res)

}
