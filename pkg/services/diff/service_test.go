// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package diff

import (
	"errors"
	"testing"

	"github.com/osrgroup/product-model-toolkit/pkg/db/memory"
)

func TestFindAllProducts(t *testing.T) {
	r := inMemRepo()
	s := NewService(r)

	prods, err := s.FindAllProducts()
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	if len(prods) != 2 {
		t.Errorf("Expected %v products as result, but go %v", 2, len(prods))
	}
}

func TestFindProdcutByID(t *testing.T) {
	r := inMemRepo()
	s := NewService(r)

	prod, err := s.FindProductByID(2)
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	if prod.ID != 2 {
		t.Errorf("Expected found product to have ID %v , but go %v", 2, prod.ID)
	}
}

func TestFindProdcutByID_ErrNotFound(t *testing.T) {
	r := inMemRepo()
	s := NewService(r)

	_, err := s.FindProductByID(-1)
	if errors.Is(err, ErrNotFound) {
		t.Errorf("Expected error to be '%v', but got '%v'", ErrNotFound, err)
	}
}

func inMemRepo() *memory.DB {
	repo := new(memory.DB)
	repo.AddSampleData()
	return repo
}
