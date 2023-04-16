package main

import (
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Println("Prime was induced")

	n := in.N
	i := uint32(2)

	for n > 1 {
		if n%i == 0 {
			n = n / i
			stream.Send(&pb.PrimeResponse{Res: i})
		} else {
			i++
		}
	}

	return nil
}
