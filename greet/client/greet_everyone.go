package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Abhinav2408/grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Abhinav"},
		{FirstName: "Kushagra"},
		{FirstName: "Rohit"},
	}

	//we will use go routines

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println("Sending request " + req.FirstName)
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

			log.Println(res.Result)
		}
		close(waitc)
	}()

	<-waitc
	//this ensures that close waitc is achieved
}
