// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package hasher

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert"
)

// Hasher represents a File-Hasher converter
type Hasher struct{}

type artifact struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Hash  hash   `json:"hashes"`
}

type hash struct {
	MD5    string `json:"md5,omitempty"`
	SHA1   string `json:"sha1,omitempty"`
	SHA256 string `json:"sha256,omitempty"`
}

// Convert converts a File-Hasher representation into a Product Model representation.
func (Hasher) Convert(input io.Reader) (*model.Product, error) {
	byteInput := new(bytes.Buffer)
	byteInput.ReadFrom(input)

	body := convert.TrimUTF8prefix(byteInput.Bytes())
	var result []artifact
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func asProductModel(artifacts []artifact) (*model.Product, error) {
	if len(artifacts) < 1 {
		return nil, errors.New("Artifact array should have at least one element")
	}

	var comps []model.Component = make([]model.Component, 0, len(artifacts))

	for _, artifact := range artifacts[:] {
		comp := asComponent(artifact)
		comps = append(comps, comp)

	}

	return &model.Product{
		Name:       "new Product",
		Components: comps,
	}, nil
}

func asComponent(art artifact) model.Component {
	return model.Component{
		Pkg:  art.Path,
		Name: art.Name,
	}
}
