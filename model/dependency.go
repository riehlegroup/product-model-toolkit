// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import "fmt"

// DepGraph represents a component dependency graph.
type DepGraph struct {
	Deps map[CmpID]map[CmpID]Dependency `json:"dependencies, omitempty"`
}

// Dependency represents the relationship between two components.
type Dependency struct {
	From    CmpID
	To      CmpID
	Linking LinkingType
}

// LinkingType represents the type of linking between dependencies, e.g. static linking.
type LinkingType string

// NewDepGraph create a new component dependency graph.
func NewDepGraph() *DepGraph {
	return &DepGraph{
		Deps: make(map[CmpID]map[CmpID]Dependency),
	}
}

const (
	// StaticLinked represents static linking between components
	StaticLinked LinkingType = "STATIC_LINKED"
	// DynamicLinked represents dynamic linking between components
	DynamicLinked LinkingType = "DYNAMIC_LINKED"
)

// String returns a string representation of a dependency.
func (d *Dependency) String() string {
	return fmt.Sprintf("(%s) -> (%s) [linking: '%s']", d.From, d.To, d.Linking)
}

// AddDependency adds a new dependency between two components.
func (g *DepGraph) AddDependency(from, to CmpID, linking LinkingType) {
	if _, ok := g.Deps[from]; !ok {
		g.Deps[from] = make(map[CmpID]Dependency)
	}

	g.Deps[from][to] = Dependency{From: from, To: to, Linking: linking}
}
