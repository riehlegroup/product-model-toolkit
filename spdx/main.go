package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	pb "pmt/model"
)

const (
	port = ":56985"
)

const (
	defaultHost = "datastore:27017"
)




func main() {
	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register reflection service on gRPC server.
	reflection.Register(s)

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer client.Disconnect(context.Background())

	productCollection := client.Database("pmt").Collection("products")
	repository := &MongoRepository{productCollection}

	pb.RegisterBomServiceServer(s, &handler{repository})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}