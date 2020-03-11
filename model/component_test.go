package model

import (
	"testing"
)

func TestID(t *testing.T) {
	c := &Component{Name: "Product", Pkg: "org.pmt.model", Version: "1.2.3-beta"}

	result := string(c.ID())

	expect := "org.pmt.model:Product:1.2.3-beta"

	if result != expect {
		t.Errorf("Expected component ID to be '%v', but got '%v'.", expect, result)
	}
}
