// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"errors"
	"math/rand"
	"fmt"
	"github.com/osrgroup/product-model-toolkit/model"
)

var (
	// ErrNotFound if a entity couldn't be found in the db.
	ErrNotFound = errors.New("entity not found")
	// ErrAlreadyExist if a entity with the same ID already exist in the db.
	ErrAlreadyExist = errors.New("entity already exist")
)

// DB represents a database.
type DB struct {
	products []model.Product
}

// FindAllProducts returns all Products from the DB.
func (db *DB) FindAllProducts() ([]model.Product, error) {
	return db.products, nil
}

// FindProductByID return the Product with the given ID from the DB.
func (db *DB) FindProductByID(id int) (model.Product, error) {
	for i := range db.products {
		if db.products[i].ID == id {
			return db.products[i], nil
		}
	}

	return model.Product{}, ErrNotFound
}

// SaveProduct store a Prodact to the DB.
func (db *DB) SaveProduct(prod *model.Product) (model.Product, error) {
	if prod.ID == 0 {
		prod.ID = rand.Intn(2147483647)
	}

	found := db.productExists(prod.ID)
	if !found {
		db.products = append(db.products, *prod)

		return *prod, nil
	}

	return model.Product{}, ErrAlreadyExist
}

func (db *DB) productExists(id int) bool {
	for i := range db.products {
		if db.products[i].ID == id {
			return true
		}
	}

	return false
}

// AddSampleData adds dummy data to the DB.
func (db *DB) AddSampleData() {
	prod1 := &model.Product{
		ID:      1,
		Name:    "Product 1",
		Version: "1.0.0",
		VCS:     "github.com/prod1",
	}

	prod2 := &model.Product{
		ID:      2,
		Name:    "Product 2",
		Version: "2.0.0",
		VCS:     "github.com/prod2",
	}

	_, err := db.SaveProduct(prod1)
	if err != nil {fmt.Println(err)}
	_, err = db.SaveProduct(prod2)
	if err != nil {fmt.Println(err)}
}
