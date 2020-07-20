// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"testing"
)

func TestAddRelation(t *testing.T) {
	g := NewDepGraph()

	g.AddDependency("a", "b", StaticLinked)
	g.AddDependency("a", "c", StaticLinked)
	g.AddDependency("c", "d", DynamicLinked)

	a := g.Deps["a"]
	if len(a) != 2 {
		t.Errorf("Expected amount of dependencies from a to be %v, but got %v", 2, len(a))
	}

	c := g.Deps["c"]
	if len(c) != 1 {
		t.Errorf("Expected amount of dependencies from c to be %v, but got %v", 1, len(c))
	}
}

func TestString(t *testing.T) {
	g := NewDepGraph()
	g.AddDependency("a", "b", StaticLinked)
	g.AddDependency("c", "d", DynamicLinked)

	a := g.Deps["a"]["b"]
	aShould := "(a) -> (b) [linking: 'STATIC_LINKED']"
	if a.String() != aShould {
		t.Errorf("Expected a.String to be '%v', but got '%v'", aShould, a.String())
	}

	c := g.Deps["c"]["d"]
	cShould := "(c) -> (d) [linking: 'DYNAMIC_LINKED']"
	if a.String() != aShould {
		t.Errorf("Expected a.String to be '%v', but got '%v'", cShould, c.String())
	}
}
