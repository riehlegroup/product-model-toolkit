package server

import (
	"github.com/osrgroup/product-model-toolkit/pkg/server/dgraph"
	"github.com/osrgroup/product-model-toolkit/pkg/server/http/rest"
)

type Instance struct {
	REST *rest.Instance
	DB   *dgraph.DB
}
