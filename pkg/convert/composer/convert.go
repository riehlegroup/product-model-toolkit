// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package composer

import (
	"encoding/json"
	"strings"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/convert"
)

// Composer represents a PHP Composer converter
type Composer struct{}

type composerDocComp struct {
	Name        string
	Version     string
	Description string
	License     []string
}

type composerDoc struct {
	Installed []composerDocComp
}

// Convert converts a PHP Composer representation into a Product Model representation.
func (Composer) Convert(doc []byte) (*model.Product, error) {
	body := convert.TrimUTF8prefix(doc)
	var result composerDoc
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		Name:       "unkown",
		Components: extractComponents(&result),
	}, nil
}

func extractComponents(input *composerDoc) []model.Component {
	var comps []model.Component
	for _, c := range input.Installed {
		licenses := strings.Join(c.License, ", ")
		comp := &model.Component{
			Name:    c.Name,
			Version: c.Version,
			License: licenses,
		}
		comps = append(comps, *comp)
	}

	return comps
}
