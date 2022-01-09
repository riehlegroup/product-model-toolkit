// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"testing"

	"github.com/osrgroup/product-model-toolkit/model"
)

var p = &model.Product{ID: 1, Name: "Prod A"}

func TestFindAllProducts(t *testing.T) {
	db := new(DB)

	prods, _ := db.FindAllProducts()
	size := len(prods)
	if size != 0 {
		t.Errorf("Expected product array size to be 0, but got %v", size)
	}
}

func TestFindAllProductsWithSampleData(t *testing.T) {
	db := new(DB)
	db.AddSampleData()

	prods, _ := db.FindAllProducts()
	size := len(prods)
	if size != 2 {
		t.Errorf("Expected product array size to be 2, but got %v", size)
	}
}

func TestFindProductByID(t *testing.T) {
	db := new(DB)
	db.AddSampleData()

	prod, err := db.FindProductByID(2)
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	if prod.ID != 2 {
		t.Errorf("Expected product ID to be 2, but got %v", prod.ID)
	}
}

func TestFindProductByIDNotFound(t *testing.T) {
	db := new(DB)
	db.AddSampleData()

	_, err := db.FindProductByID(3)
	if err.Error() != ErrNotFound.Error() {
		t.Errorf("Expected error to be ErrNotFound, but got %v", err)
	}
}

func TestSaveProduct(t *testing.T) {
	db := new(DB)

	_, err := db.SaveProduct(p)
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	size := len(db.products)
	if size != 1 {
		t.Errorf("Expected size to be 0, but got %v", size)
	}
	pResult := db.products[0]
	if pResult.ID != p.ID {
		t.Errorf("Expected products to have the same ID, but got %v", pResult.ID)
	}
}

func TestSaveProductAlreadyExist(t *testing.T) {
	db := new(DB)

	db.SaveProduct(p)
	_, err := db.SaveProduct(p)

	if err != ErrAlreadyExist {
		t.Errorf("Expected error ErrAlreadyExist, but got %v", err)
	}

	size := len(db.products)
	if size != 1 {
		t.Errorf("Expected size to be 0, but got %v", size)
	}
}
