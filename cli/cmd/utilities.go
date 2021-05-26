package cmd

import (
	"google.golang.org/grpc"
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
		return nil, err
	}
	bomServiceClient := pb.NewBomServiceClient(conn)
	client := &Client{
		bomServiceClient,
		conn,
	}
	return client, nil
}

