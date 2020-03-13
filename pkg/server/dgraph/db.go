package dgraph

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// DB represents a Dgraph db connection.
type DB struct {
	client *dgo.Dgraph
}

// DefaultURI is the default URI string for a local Dgraph server.
const DefaultURI string = "localhost:9080"

// NewClient returns a DB struct with an initialized Dgraph client.
func NewClient(URI string) *DB {
	d, err := grpc.Dial(URI, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)

	log.Println("[DB] Created new DB client for URI: ", URI)
	log.Println("[DB] Connection state: ", d.GetState().String())

	return &DB{
		client: c,
	}
}

// DropAll drops all elements in the db.
func (db *DB) DropAll() error {
	log.Println("[DB] Drop all elements.")
	err := db.client.Alter(
		context.Background(),
		&api.Operation{DropOp: api.Operation_ALL})

	return err
}
