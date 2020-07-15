// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package importing

import (
	"bytes"
	"os"
	"testing"
)

const testFile = "convert/composer/test/example.json"

func TestComposerRead(t *testing.T) {
	jsonFile, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Unable to read %s to start tests", testFile)
	}
	defer jsonFile.Close()

	s := NewService()
	p, err := s.ComposerRead(jsonFile)
	if err != nil {
		t.Errorf("Expected to read composer file without exception, but got %v", err)
	}

	if p == nil {
		t.Error("Expected product to be not nil")
	}

	if p.Name != "new Product" {
		t.Errorf("Expected product to have name 'new Product', but got '%v'", p.Name)
	}

	if len(p.Components) != 386 {
		t.Errorf("Expected amount of components to be %v, but got %v", 386, len(p.Components))
	}
}

func TestComposerRead_Empty(t *testing.T) {
	s := NewService()

	_, err := s.ComposerRead(bytes.NewReader(nil))
	if err == nil {
		t.Error("Expected ComposerRead() to return error for empty input")
	}
}
