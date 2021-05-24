package cmd

import (
	"google.golang.org/grpc"
	"log"
	pb "pmt/model"
)

type Client struct {
	pb.BomServiceClient
	*grpc.ClientConn
}

func createGrpcClient() (*Client, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:56985", grpc.WithInsecure()) // TODO(change the hard coded address)
	if err != nil {
		log.Printf("Did not connect: %v\n", err)
		return nil, err
	}

	client := pb.NewBomServiceClient(conn)
	bomClient := &Client{
		client,
		conn,
	}
	return bomClient, nil
}

