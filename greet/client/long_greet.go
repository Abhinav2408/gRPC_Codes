package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Abhinav2408/grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Abhinav"},
		{FirstName: "Kushagra"},
		{FirstName: "Avi"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Result)

}
