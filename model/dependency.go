// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

// Dependency represents the relationship between two components.
type Dependency struct {
	Type DepType
	From *Component
	To   *Component
}

// Next returns true if dependency points to a component.
func (d *Dependency) Next() bool {
	return d.To != nil
}

// DepType represents the type of a dependency, e.g. static linking.
type DepType string

const (
	StaticLinked  DepType = "STATIC_LINKED"
	DynamicLinked DepType = "DYNAMIC_LINKED"
)

type Iterator interface {
	Next() bool
}
