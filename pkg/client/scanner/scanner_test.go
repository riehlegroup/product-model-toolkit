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
	licensee := FromStr("licensee")
	scancode := FromStr("scancode")

	if licensee.String() != Licensee.String() {
		t.Errorf("Expected returned tool to be %v, but got %v", Licensee.String(), Licensee.String())
	}

	if scancode.String() != Scancode.String() {
		t.Errorf("Expected returned tool to be %v, but got %v", Scancode.String(), scancode.String())
	}
}

func TestFromStr_CaseInsensitive(t *testing.T) {
	licensee := FromStr("lIceNsEE")
	scancode := FromStr("scANcoDE")

	if licensee.String() != Licensee.String() {
		t.Error("Expected input to work case insensitive")
	}

	if scancode.String() != Scancode.String() {
		t.Error("Expected input to work case insensitive")
	}
}

func TestFromStr_DefaultCase(t *testing.T) {
	empty := FromStr("")
	unknown := FromStr("unknown tool")

	if empty.String() != Default.String() {
		t.Errorf("Expected tool from empty input to be %v, but got %v", Default.String(), empty.String())
	}

	if unknown.String() != Default.String() {
		t.Errorf("Expected tool from unknown input to be %v, but got %v", Default.String(), unknown.String())
	}

}

func TestAvailable_NotEmpty(t *testing.T) {
	size := len(Available)

	if size != 3 {
		t.Errorf("Expected size to be 2, but got %v", size)
	}
}
