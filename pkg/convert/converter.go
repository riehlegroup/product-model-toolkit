// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package convert

import (
	"bytes"
	"io"

	"github.com/osrgroup/product-model-toolkit/model"
)

// Converter is the interface all converter implementations need to fulfill.
type Converter interface {
	// Convert converts a doc to the product model representation.
	Convert(io.Reader) (*model.Product, error)
}

// TrimUTF8prefix returns doc without UTF8 prefix string
func TrimUTF8prefix(doc []byte) []byte {
	return bytes.TrimPrefix(doc, []byte("\xef\xbb\xbf"))
}
