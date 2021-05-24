package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvloader"
	"io"
	"io/ioutil"
	pb "pmt/model"
)

type handler struct {
	//mu sync.RWMutex
	repository
}

// SPDX
func AsSPDX(input io.Reader) (*spdx.Document2_1, error) {
	doc, err := tvloader.Load2_1(input)
	if err != nil {
		msg := fmt.Sprintf("error while parsing SPDX body: %v", err)
		return nil, errors.New(msg)
	}

	return doc, nil
}

func (h *handler) CreateBom(ctx context.Context, req *pb.InputValue) (*pb.Response, error) {

	data, err := ioutil.ReadFile(req.FilePath)
	if err != nil {
		return nil, err
	}

	product := &pb.Product{}
	if err := proto.Unmarshal(data, product); err != nil {
		err := errors.New(fmt.Sprintf("failed to parse product: %v", err))
		return nil, err
	}

	// do the logic here (creating BoM!
	switch req.Type {
	case pb.InputType_SPDX:
		ioReaderData := bytes.NewReader(data)
		doc, err := AsSPDX(ioReaderData)
		if err != nil {
			return nil, err
		}
		fmt.Println(doc)
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
