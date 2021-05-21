package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	bomPb "pmt/bom/proto"
	productPb "pmt/product/proto"
	"sync"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*bomPb.InputValue) (*productPb.Product, error)
}

type Repository struct {
	mu sync.RWMutex
	products []*productPb.Product
}

// Create a new bom
func (repo *Repository) Create(inputValue *bomPb.InputValue) (*productPb.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	switch inputValue.Type {
	case bomPb.InputType_SPDX:
		fmt.Println("spdx")
	case bomPb.InputType_HUMAN:
		fmt.Println("human")
	case bomPb.InputType_CUSTOM:
		fmt.Println("custom")
	}

	return &productPb.Product{}, nil
}

type service struct {
	repo repository
}

func (s *service) CreateBom(ctx context.Context, req *bomPb.InputValue) (*bomPb.Response, error) {

	product, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &bomPb.Response{Created: true, Product: product}, nil
}

func main() {

	repo := &Repository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	bomPb.RegisterBomServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}