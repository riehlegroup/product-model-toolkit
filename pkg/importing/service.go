// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package importing

import (
	"fmt"
	"io"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert/composer"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert/hasher"
	"github.com/pkg/errors"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvloader"
)

// Service  provides BOM import operations.
type Service interface {
	ComposerImport(io.Reader) (*model.Product, error)
	SPDX(io.Reader) (*spdx.Document2_1, error)
	FileHasherImport(io.Reader) (*model.Product, error)
}

type service struct {
	r repository
}

type repository interface {
	SaveProduct(prod *model.Product) (model.Product, error)
}

// NewService creates a querying service with all necessary dependencies.
func NewService(r repository) Service {
	return &service{r}
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
