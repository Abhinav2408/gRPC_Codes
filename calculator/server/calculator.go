package main

import (
	"context"
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator Function was invoked with %v\n", in)

	return &pb.CalculatorResponse{
		Result: in.One + in.Two,
	}, nil
}
