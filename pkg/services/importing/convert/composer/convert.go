// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package composer

import (
	"bytes"
	"encoding/json"
	convert2 "github.com/osrgroup/product-model-toolkit/pkg/services/importing/convert"
	"io"
	"strings"

	"github.com/osrgroup/product-model-toolkit/model"
)

// Composer represents a PHP Composer converter
type Composer struct{}

type composerDocComp struct {
	Name        string
	Version     string
	Description string
	License     []string
	Requires    []composerDocComp
}

type composerDoc struct {
	Installed []composerDocComp
}

type mapComp map[model.CmpID]model.Component

// Convert converts a PHP Composer representation into a Product Model representation.
func (Composer) Convert(input io.Reader) (*model.Product, error) {

	byteInput := new(bytes.Buffer)
	byteInput.ReadFrom(input)

	body := convert2.TrimUTF8prefix(byteInput.Bytes())
	var result composerDoc
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	compAsMap := make(mapComp, 2*len(result.Installed))
	extractDependencies(&result.Installed, compAsMap)
	comps := compMapToSlice(compAsMap)

	return &model.Product{
		Name:       "new Product",
		Components: comps,
	}, nil
}

func compMapToSlice(compAsMap mapComp) []model.Component {
	compSlice := make([]model.Component, 0, len(compAsMap))
	for _, value := range compAsMap {
		compSlice = append(compSlice, value)
	}
	return compSlice
}

func extractDependencies(input *[]composerDocComp, comps map[model.CmpID]model.Component) {
	for _, c := range *input {
		licenses := strings.Join(c.License, ", ")
		comp := model.Component{
			Name:    c.Name,
			Version: c.Version,
			License: model.License{DeclaredLicense: licenses},
		}

		_, ok := comps[comp.ID()]
		if !ok {
			comps[comp.ID()] = comp
		}

		if len(c.Requires) > 0 {
			extractDependencies(&c.Requires, comps)
		}
	}
}
