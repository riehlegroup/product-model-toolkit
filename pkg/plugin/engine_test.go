// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"testing"
)

func TestString(t *testing.T) {
	plugin1 := &Plugin{Name: "myName", Version: "0.0.0"}
	expected1 := "myName (0.0.0)"

	plugin2 := &Plugin{Name: "otherName", Version: "2.2.2-beta"}
	expected2 := "otherName (2.2.2-beta)"

	if plugin1.String() != expected1 {
		t.Errorf("Expected string representation to be '%s', but got '%s'", expected1, plugin1.String())
	}

	if plugin2.String() != expected2 {
		t.Errorf("Expected string representation to be '%s', but got '%s'", expected2, plugin2.String())
	}
}
