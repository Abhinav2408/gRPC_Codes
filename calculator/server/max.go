package main

import (
	"io"
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max invoked")

	var res int32 = -1e5

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		if req.Num > res {
			res = req.Num
		}
		err = stream.Send(&pb.MaxResponse{
			Res: res,
		})

		if err != nil {
			log.Fatalln(err)
		}

	}
}
