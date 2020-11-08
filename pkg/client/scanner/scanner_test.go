// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

import "testing"

func TestString(t *testing.T) {
	tool1 := &Tool{Name: "myName", Version: "0.0.0"}
	expected1 := "myName (0.0.0)"

	tool2 := &Tool{Name: "otherName", Version: "2.2.2-beta"}
	expected2 := "otherName (2.2.2-beta)"

	if tool1.String() != expected1 {
		t.Errorf("Expected String representation to be '%s', but got '%s'", expected1, tool1.String())
	}

	if tool2.String() != expected2 {
		t.Errorf("Expected String representation to be '%s', but got '%s'", expected2, tool2.String())
	}
}

func TestFromStr(t *testing.T) {
	licensee, found := FromStr("licensee")
	scancode, found := FromStr("scancode")

	if found == false {
		t.Errorf("Expected returned tool to be %v, but got no tool", "Licensee")
	} else if licensee.Name != "Licensee" {
		t.Errorf("Expected returned tool to be %v, but got %v", "Licensee", licensee.Name)
	}

	if found == false {
		t.Errorf("Expected returned tool to be %v, but got no tool", "Scancode")
	} else if scancode.Name != "Scancode" {
		t.Errorf("Expected returned tool to be %v, but got %v", "Scancode", scancode.Name)
	}
}

func TestFromStr_CaseInsensitive(t *testing.T) {
	licensee, found := FromStr("lIceNsEE")
	scancode, found := FromStr("scANcoDE")

	if found == false || licensee.Name != "Licensee" {
		t.Error("Expected input to work case insensitive")
	}

	if found == false || scancode.Name != "Scancode" {
		t.Error("Expected input to work case insensitive")
	}
}

func TestFromStr_DefaultCase(t *testing.T) {
	empty, found := FromStr("")
	unknown, found := FromStr("unknown tool")

	if found == true {
		t.Errorf("Expected no tool from empty input, but got %v", empty.String())
	}

	if found == true {
		t.Errorf("Expected no tool from unknown input, but got %v", unknown.String())
	}
}

func TestAvailable_NotEmpty(t *testing.T) {
	size := len(Available)

	if size != 4 {
		t.Errorf("Expected size to be 4, but got %v", size)
	}
}
