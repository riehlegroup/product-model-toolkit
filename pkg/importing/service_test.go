// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package importing

import (
	"bytes"
	"os"
	"testing"

	"github.com/osrgroup/product-model-toolkit/pkg/db/memory"
)

func TestComposerRead(t *testing.T) {
	const testFile = "convert/composer/test/example.json"
	jsonFile, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Unable to read %s to start tests", testFile)
	}
	defer jsonFile.Close()

	repo := new(memory.DB)
	s := NewService(repo)
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
	repo := new(memory.DB)
	s := NewService(repo)

	_, err := s.ComposerRead(bytes.NewReader(nil))
	if err == nil {
		t.Error("Expected ComposerRead() to return error for empty input")
	}
}

func TestFileHasherImport(t *testing.T) {
	const testFile = "convert/hasher/test/example.json"
	jsonFile, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Unable to read %s to start tests", testFile)
	}
	defer jsonFile.Close()

	repo := new(memory.DB)
	s := NewService(repo)
	p, err := s.FileHasherImport(jsonFile)
	if err != nil {
		t.Errorf("Expected to import file-hasher json without exception, but got %v", err)
	}

	if p == nil {
		t.Error("Expected product to be not nil")
	}

	if p.Name != "new Product" {
		t.Errorf("Expected product to have name 'new Product', but got '%v'", p.Name)
	}

	if len(p.Components) != 18 {
		t.Errorf("Expected amount of components to be %v, but got %v", 18, len(p.Components))
	}
}

func TestFileHasherImport_Empty(t *testing.T) {
	repo := new(memory.DB)
	s := NewService(repo)

	_, err := s.FileHasherImport(bytes.NewReader(nil))
	if err == nil {
		t.Error("Expected FileHasherImport() to return error for empty input")
	}
}
