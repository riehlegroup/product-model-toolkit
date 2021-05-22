package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "pmt/model"
	"sync"
)

const (
	port = ":56985"
)

type repository interface {
	Create(*pb.InputValue) (*pb.Product, error)
}

type Repository struct {
	mu sync.RWMutex
	products []*pb.Product
}

// Create a new bom
func (repo *Repository) Create(inputValue *pb.InputValue) (*pb.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	switch inputValue.Type {
	case pb.InputType_SPDX:
		fmt.Println("spdx")
	case pb.InputType_HUMAN:
		fmt.Println("human")
	case pb.InputType_CUSTOM:
		fmt.Println("custom")
	}

	return &pb.Product{}, nil
}

type service struct {
	repo repository
}

func (s *service) CreateBom(ctx context.Context, req *pb.InputValue) (*pb.Response, error) {
	product, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Product: product}, nil
}

func main() {

	repo := &Repository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterBomServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}