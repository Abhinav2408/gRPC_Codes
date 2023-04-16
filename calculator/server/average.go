package main

import (
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average invoked")

	var x float64 = 0.0
	var i int = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Res: x / float64(i),
			})
		}

		if err != nil {
			log.Fatal("Error occured")
		}

		x += float64(req.Num)
		i += 1
	}
}
