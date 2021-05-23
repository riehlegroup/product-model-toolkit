package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	pb "pmt/model"
	"sync"
	proto "github.com/golang/protobuf/proto"
)

type handler struct {
	mu sync.RWMutex
	repository
}

func (h *handler) CreateBom(ctx context.Context, req *pb.InputValue) (*pb.Response, error) {

	data, err := ioutil.ReadFile(req.FilePath)
	if err != nil {
		return nil, err
	}

	product := &pb.Product{}
	if err := proto.Unmarshal(data, product); err != nil {
		err := errors.New(fmt.Sprintf("Failed to parse product: %v", err))
		return nil, err
	}

	// do the logic here (creating BoM!
	switch req.Type {
	case pb.InputType_SPDX:
		//
		fmt.Println("spdx")
	case pb.InputType_HUMAN:
		//
		fmt.Println("human")
	case pb.InputType_CUSTOM:
		//
		fmt.Println("custom")
	}

	// Save our product
	if err := h.repository.Create(ctx, MarshalProduct(product)); err != nil {
		return nil, err
	}

	return &pb.Response{Result: nil, Product: product}, nil
}

func (h *handler) GetProducts(ctx context.Context, req *pb.GerRequest) (*pb.Response, error) {
	products, err := h.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := &pb.Response{
		Result:  nil,
		Products: UnmarshalProductCollection(products),
	}
	return response, nil
}
