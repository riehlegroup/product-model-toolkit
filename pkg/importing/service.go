// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package importing

import (
	"errors"
	"fmt"
	"io"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert/composer"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvloader"
)

// Service  provides BOM import operations.
type Service interface {
	Composer(io.Reader) (*model.Product, error)
	SPDX(io.Reader) (*spdx.Document2_1, error)
}

type service struct{}

// NewService creates a querying service with all necessary dependencies.
func NewService() Service {
	return &service{}
}

// Composer import a Composer representation of the BOM.
func (service) Composer(input io.Reader) (*model.Product, error) {
	var c convert.Converter = new(composer.Composer)
	prod, err := c.Convert(input)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing Composer JSON body: %v", err)
		return nil, errors.New(msg)
	}

	return prod, nil
}

// SPDX import a SPDX representation of the BOM.
func (service) SPDX(input io.Reader) (*spdx.Document2_1, error) {
	doc, err := tvloader.Load2_1(input)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing SPDX body: %v", err)
		return nil, errors.New(msg)
	}

	return doc, nil
}
