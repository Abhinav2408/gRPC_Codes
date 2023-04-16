package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	reqs := []*pb.MaxRequest{
		{Num: -2},
		{Num: 3},
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 10},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request %d\n", req.Num)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
				break
			}

			log.Printf("Max is %v\n", res)
		}

		close(waitc)
	}()

	<-waitc
}
