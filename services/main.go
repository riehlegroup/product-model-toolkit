package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"pmt/internal/cnst"
	pb "pmt/model"
)


func main() {

	// listen to the default gRPC port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cnst.DefaultGrpcPort))

	// check error of listening
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// define new gRPC server
	s := grpc.NewServer()

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// get the db host environment variable
	uri := os.Getenv("DB_HOST")

	// if the variable was empty set the default value
	if cnst.Debug {
		if uri == cnst.Empty {
			uri = cnst.MongoDBDevelopmentHost
		}
	} else {
		if uri == cnst.Empty {
			uri = cnst.MongoDBDefaultHost
		}
	}

	// define new MongoDB client
	mongoDBClient, err := createMongoDbClient(context.Background(), uri, cnst.MongoDBDefaultRetryNumber)

	// check error of creating new Mongo DB client
	if err != nil {
		log.Fatalf("db unable to connect: %v", err)
	}

	// close connection in defer
	defer mongoDBClient.Disconnect(context.Background())

	// define new product collection
	productCollection := mongoDBClient.
		Database(cnst.MongoDBDefaultDatabaseName).
		Collection(cnst.MongoDBDefaultCollectionName)

	// define new repository
	repository := &MongoRepository{productCollection}

	// register the gRPC servers of services
	pb.RegisterImportServiceServer(s, &handler{repository})

	// print out the status of gRPC server
	log.Println("Running on port:", cnst.DefaultGrpcPort)

	// check error of server running
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}