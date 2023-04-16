package main

import (
	"log"

	pb "github.com/Abhinav2408/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//second paramter is for security
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	// doCalculator(c, 5, 7)
	// doPrime(c)
	// doAverage(c)
	doMax(c)
}
