// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package hasher

import (
	"bytes"
	"encoding/json"
	"errors"
	convert2 "github.com/osrgroup/product-model-toolkit/pkg/services/importing/convert"
	"io"
	"strings"

	"github.com/osrgroup/product-model-toolkit/model"
)

// Hasher represents a File-Hasher converter
type Hasher struct{}

// Convert converts a File-Hasher representation into a Product Model representation.
func (Hasher) Convert(input io.Reader) (*model.Product, error) {
	byteInput := new(bytes.Buffer)
	byteInput.ReadFrom(input)

	body := convert2.TrimUTF8prefix(byteInput.Bytes())
	var result []model.Artifact
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	prod, err := asProductModel(result)
	if err != nil {
		return nil, err
	}

	return prod, nil
}

func asProductModel(artifacts []model.Artifact) (*model.Product, error) {
	if len(artifacts) < 1 {
		return nil, errors.New("Artifact array should have at least one element")
	}

	basePath := artifacts[0].Path

	var comps []model.Component = make([]model.Component, 0, len(artifacts))
	for _, artifact := range artifacts[:] {
		artifact.Path = removeBasePath(artifact.Path, basePath)
		comp := asComponent(artifact)
		comps = append(comps, comp)

	}

	return &model.Product{
		Name:       "new Product",
		Components: comps,
	}, nil
}

func asComponent(art model.Artifact) model.Component {
	return model.Component{
		Pkg:      art.Path,
		Name:     art.Name,
		Artifact: art,
	}
}

// removeBasePath returns the path without the leading base path.
func removeBasePath(path string, basePath string) string {
	if basePath == "/" {
		return path
	}

	return strings.TrimPrefix(path, basePath)
}
