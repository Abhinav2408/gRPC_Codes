package main

import (
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalln(err)
		}
	}

}
