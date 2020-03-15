package dgraph

import "github.com/osrgroup/product-model-toolkit/model"

// DProduct extends model.Product with fields only required by Dgraph.
type DProduct struct {
	model.Product
	DType []string `json:"dgraph.type,omitempty"`
}

// DComponent extends model.Component with fields only required by Dgraph.
type DComponent struct {
	model.Component
	DType []string `json:"dgraph.type,omitempty"`
}
