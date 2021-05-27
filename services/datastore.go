package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func createMongoDbClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	// create a MongoDB connection
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// ping the connection and check the error
	//if err := conn.Ping(ctx, nil); err != nil {
	if err != nil {
		if retry >= 3 {
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return createMongoDbClient(ctx, uri, retry)
	}

	return conn, err
}