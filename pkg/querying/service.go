// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package querying

import (
	"errors"

	"github.com/osrgroup/product-model-toolkit/model"
)

var (
	// ErrNotFound if a entity couldn't be found in the db.
	ErrNotFound = errors.New("Entity not found")
	// ErrAlreadyExist if a entity with the same ID already exist in the db.
	ErrAlreadyExist = errors.New("Entity already exist")
)

// Repository provides access to the product db.
type Repository interface {
	// FindAllProducts returns a list of all products saved in db.
	FindAllProducts() (*[]model.Product, error)
	// FindProductByID returns the product with the given ID.
	FindProductByID(id string) (*model.Product, error)
}

// Service  provides product querying operations.
type Service interface {
	FindAllProducts() (*[]model.Product, error)
	FindProductByID(id string) (*model.Product, error)
}

type service struct {
	r Repository
}

// NewService creates a querying service with all necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// FindAllProducts returns all existing products.
func (s *service) FindAllProducts() (*[]model.Product, error) {
	return s.r.FindAllProducts()
}

// FindProductByID returns the product with the given ID.
func (s *service) FindProductByID(id string) (*model.Product, error) {
	return s.r.FindProductByID(id)
}
