package model

import (
	"testing"
)

var (
	cmp1 *Component = &Component{Pkg: "com.test", Name: "Cmp1", Version: "1.0.0"}
	cmp2 *Component = &Component{Pkg: "com.test", Name: "Cmp2", Version: "2.0.0"}
)

func TestNextTrue(t *testing.T) {
	d := &Dependency{Type: StaticLinked, From: cmp1, To: cmp2}

	if !d.Next() {
		t.Error("Expected dependency to have next element")
	}
}

func TestNextFalse(t *testing.T) {
	d := &Dependency{Type: StaticLinked, From: cmp1, To: nil}

	if d.Next() {
		t.Error("Expected dependency to have no next element")
	}
}
