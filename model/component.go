// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import "fmt"

// Component represents a unit of compositon of the product, e.g. class, lib, module.
type Component struct {
	UID       string `json:"id"`
	Name      string `json:"name,omitempty"`
	Pkg       string `json:"package,omitempty"`
	Version   string `json:"version,omitempty"`
	License   string `json:"license,omitempty"`
	Copyright `json:"-"`
	Artifact  `json:"-"`
}

// CmpID represents an component identifier which is a combination of package + name + version.
type CmpID string

// ID returns an identifier which is a combination of package + name + version.
func (c *Component) ID() CmpID {
	id := fmt.Sprintf("%v:%v:%v", c.Pkg, c.Name, c.Version)
	return CmpID(id)
}
