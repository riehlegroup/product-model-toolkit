// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import "errors"

// UsageType represents the scenario in which a product ist used, e.g. cloud service or internal usage only.
type UsageType string

const (
	OnPremise    UsageType = "on-premise"
	CloudService           = "cloud-service"
	Library                = "library"
	Internal               = "internal"
)

func (ut UsageType) isValid() error {
	switch ut {
	case OnPremise, CloudService, Library, Internal:
		return nil
	}

	return errors.New("Invalid usage type")
}
