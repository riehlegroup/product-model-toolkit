// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package services

import (
	"fmt"
	convert "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert"
	composer "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert/composer"
	hasher "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert/hasher"
	"io"

	"github.com/pkg/errors"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvloader"
	"github.com/osrgroup/product-model-toolkit/model"
)

var (
	// ErrNotFound if a entity couldn't be found in the db.
	ErrNotFound = errors.New("entity not found")
)

// Repository provides access to the product db.
type Repository interface {
	// FindAllProducts returns a list of all products saved in db.
	FindAllProducts() ([]model.Product, error)
	// FindProductByID returns the product with the given ID.
	FindProductByID(id int) (model.Product, error)

	SaveProduct(prod *model.Product) (model.Product, error)
}

// Service  provides product querying operations.
type Service interface {
	FindAllProducts() ([]model.Product, error)
	FindProductByID(id int) (model.Product, error)

	// import
	ComposerImport(io.Reader) (*model.Product, error)
	SPDX(io.Reader) (*spdx.Document2_1, error)
	FileHasherImport(io.Reader) (*model.Product, error)
}

type service struct {
	r Repository
}

// NewService creates a querying service with all necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// FindAllProducts returns all existing products.
func (s *service) FindAllProducts() ([]model.Product, error) {
	return s.r.FindAllProducts()
}

// FindProductByID returns the product with the given ID.
func (s *service) FindProductByID(id int) (model.Product, error) {
	return s.r.FindProductByID(id)
}

// ComposerImport import a Composer representation of the BOM and store it in the DB.
func (s *service) ComposerImport(input io.Reader) (*model.Product, error) {
	var c convert.Converter = new(composer.Composer)
	prod, err := c.Convert(input)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing Composer JSON body: %v", err)
		return nil, errors.New(msg)
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		msg := fmt.Sprintf("Error while saving the product to the DB: %v", err)
		return nil, errors.New(msg)
	}

	return &pSaved, nil
}

// SPDX import a SPDX representation of the BOM.
func (s *service) SPDX(input io.Reader) (*spdx.Document2_1, error) {
	doc, err := tvloader.Load2_1(input)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing SPDX body: %v", err)
		return nil, errors.New(msg)
	}

	return doc, nil
}

// FileHasherImport import a File-Hasher representation of the BOM and store it in the DB.
func (s *service) FileHasherImport(input io.Reader) (*model.Product, error) {
	var c convert.Converter = new(hasher.Hasher)

	prod, err := c.Convert(input)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing File-Hasher body")
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		return nil, errors.Wrap(err, "Error while saving the product to the DB")
	}

	return &pSaved, nil
}
