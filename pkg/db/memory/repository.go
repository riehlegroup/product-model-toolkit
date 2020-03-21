package memory

import (
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
)

// DB represents a database.
type DB struct {
	products []model.Product
}

// FindAllProducts returns all Products from the DB.
func (db *DB) FindAllProducts() (*[]model.Product, error) {
	return &db.products, nil
}

// FindProductByID return the Product with the given ID from the DB.
func (db *DB) FindProductByID(id string) (*model.Product, error) {
	for i := range db.products {
		if db.products[i].UID == id {
			return &db.products[i], nil
		}
	}

	return nil, querying.ErrNotFound
}

// SaveProduct store a Prodact to the DB.
func (db *DB) SaveProduct(prod *model.Product) error {
	found := db.productExists(prod.UID)
	if !found {
		db.products = append(db.products, *prod)
		return nil
	}

	return querying.ErrAlreadyExist
}

func (db *DB) productExists(id string) bool {
	for i := range db.products {
		if db.products[i].UID == id {
			return true
		}
	}

	return false
}

// AddSampleData adds dummy data to the DB.
func (db *DB) AddSampleData() {
	prod1 := &model.Product{
		UID:     "1",
		Name:    "Product 1",
		Version: "1.0.0",
		VCS:     "github.com/prod1",
	}

	prod2 := &model.Product{
		UID:     "2",
		Name:    "Product 2",
		Version: "2.0.0",
		VCS:     "github.com/prod2",
	}

	db.SaveProduct(prod1)
	db.SaveProduct(prod2)
}
