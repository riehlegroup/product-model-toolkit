package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "pmt/pb"
)

const (
	address         = "localhost:56985"
	defaultFilename = "product.json"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewBomServiceClient(conn)

	// Contact the server and print out its response.
	fileName := defaultFilename
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	inputValue := &pb.InputValue{
		FileName: fileName,
		Type:     pb.InputType_SPDX, // test
	}

	r, err := client.CreateBom(context.Background(), inputValue)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
}
