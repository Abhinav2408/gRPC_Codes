package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet invoked.")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("Error occured")
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
