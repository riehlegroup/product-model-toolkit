package dgraph

import "github.com/osrgroup/product-model-toolkit/model"

// DModel interface conatins methods for models used for Dgraph storage.
type DModel interface {
	DefaultDType() string
}

// DProduct extends model.Product with fields only required by Dgraph.
type DProduct struct {
	model.Product
	DType []string `json:"dgraph.type,omitempty"`
}

// DefaultDType returns string which defines the type for a Dgraph node.
func (d *DProduct) DefaultDType() string {
	return "Product"
}

// DComponent extends model.Component with fields only required by Dgraph.
type DComponent struct {
	model.Component
	DType []string `json:"dgraph.type,omitempty"`
}

// DefaultDType returns string which defines the type for a Dgraph node.
func (d *DComponent) DefaultDType() string {
	return "Component"
}
