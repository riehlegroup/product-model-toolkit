package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvloader"
	"pmt/internal/cnst"
	"io"
	"io/ioutil"
	pb "pmt/model"
)

type handler struct {
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

func (h *handler) CreateImport(ctx context.Context, req *pb.ImportInput) (*pb.Response, error) {
	// read the imported file
	data, err := ioutil.ReadFile(req.FilePath)

	// check the error of reading the file
	if err != nil {
		return nil, err
	}

	// define new product struct
	product := &pb.Product{}

	// unmarshal the data into product and check the error
	if err := proto.Unmarshal(data, product); err != nil {
		err := errors.New(fmt.Sprintf("failed to parse product: %v", err))
		return nil, err
	}

	// switch case over the type of import
	switch req.Type {
	// in case if it is spdx file
	case cnst.SPDX:
		// convert data bytes to ioReader format
		ioReaderData := bytes.NewReader(data)

		// convert the input to the document
		doc, err := AsSPDX(ioReaderData)

		// check the error of conversion
		if err != nil {
			return nil, err
		}

		fmt.Println(doc)
		fmt.Println("spdx")

	// in case if it is composer file
	case cnst.Composer:
		//
		fmt.Println("composer")

	// in case if it is file-hasher file
	case cnst.FileHasher:
		//
		fmt.Println("file hasher")
	default:
		return nil, errors.New("invalid type")
	}

	// save the product and check the error
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
		Result:   nil,
		Products: UnmarshalProductCollection(products),
	}
	return response, nil
}
