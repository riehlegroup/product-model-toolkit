// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"regexp"
	"testing"
)

func TestIsSemanticVersion(t *testing.T) {
	r, _ := regexp.Compile("^([0-9]+)\\.([0-9]+)\\.([0-9]+)(?:-([0-9A-Za-z-]+(?:\\.[0-9A-Za-z-]+)*))?(?:\\+[0-9A-Za-z-]+)?$")
	result := r.MatchString(Name())
	if !result {
		t.Errorf("Expected version to be in semantic version style, but got %s", Name())
	}
}

func TestName(t *testing.T) {
	if version != Name() {
		t.Error("Expected Name() to return version")
	}
}
