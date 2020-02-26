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
