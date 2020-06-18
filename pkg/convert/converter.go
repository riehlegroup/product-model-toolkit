// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package convert

import "github.com/osrgroup/product-model-toolkit/model"

// Converter is the interface all converter implementations need to fulfill.
type Converter interface {
	// Convert converts a doc to the product model representation.
	Convert([]byte) (*model.Product, error)
}
