// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import "errors"

// UsageType represents the scenario in which a product ist used, e.g. cloud service or internal usage only.
type UsageType string

const (
	// OnPremise represents a on-premise installation of the product.
	OnPremise UsageType = "on-premise"
	// CloudService represents a usage of the product as cloud service.
	CloudService = "cloud-service"
	// Library represents a usage as library for other products.
	Library = "library"
	// Internal represents a internal usage of the products.
	Internal = "internal"
)

func (ut UsageType) isValid() error {
	switch ut {
	case OnPremise, CloudService, Library, Internal:
		return nil
	}

	return errors.New("Invalid usage type")
}
