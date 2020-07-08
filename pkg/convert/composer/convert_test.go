// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package composer

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/convert"
)

const testFile = "test/example.json"

var exampleDoc []byte

// TestMain runs before all tests once
func TestMain(m *testing.M) {
	var err error
	exampleDoc, err = readExampleDoc()
	if err != nil {
		log.Fatalf("Unable to read %s to start tests", testFile)
		os.Exit(-1)
	}
	os.Exit(m.Run())
}

func readExampleDoc() ([]byte, error) {
	jsonFile, err := os.Open(testFile)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func TestConvert(t *testing.T) {
	var c convert.Converter = new(Composer)
	p, err := c.Convert(exampleDoc)

	t.Run("not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Expeted convert func to not return an error, but got %v", err.Error())
		}
	})

	t.Run("not nil", func(t *testing.T) {
		if p == nil {
			t.Error("Expected product to be not nil")
		}
	})

	t.Run("product name", func(t *testing.T) {
		if p.Name != "unkown" {
			t.Errorf("Expected amount of components to be %v, but got %v", "unkown", p.Name)
		}
	})

	t.Run("component amount", func(t *testing.T) {
		if len(p.Components) != 197 {
			t.Errorf("Expected amount of components to be %v, but got %v", 197, len(p.Components))
		}
	})

	t.Run("contain first component 'bluespice/about'", func(t *testing.T) {
		if !model.ContainsComp(p.Components, ":bluespice/about:dev-REL1_31") {
			t.Errorf("Expected component 'bluespice/about' to be present")
		}
	})

	t.Run("contain last component 'zordius/lightncandy'", func(t *testing.T) {
		if !model.ContainsComp(p.Components, ":zordius/lightncandy:v0.23") {
			t.Errorf("Expected component 'zordius/lightncandy' to be present")
		}
	})
}

func TestConvert_Empty(t *testing.T) {
	c := &Composer{}
	p, err := c.Convert([]byte{})
	if err == nil {
		t.Error("Expeted returning an error for empty doc")
	}
	if p != nil {
		t.Error("Expeted resulting product for empty doc to be nil")
	}
}

func TestExtractComponents(t *testing.T) {
	a := &composerDocComp{
		Name:    "CompA",
		Version: "0.4.2",
		License: []string{"MIT", "Apache-2.0"},
	}
	b := &composerDocComp{
		Name:    "CompB",
		Version: "1.1.1",
		License: []string{"GPL"},
	}
	input := &composerDoc{
		Installed: []composerDocComp{*a, *b},
	}

	comps := make(mapComp)
	extractDependencies(&input.Installed, comps)

	if len(comps) != 2 {
		t.Errorf("Expected to return %v components, but got %v", 2, len(comps))
	}

	compA, ok := comps[":CompA:0.4.2"]
	if !ok {
		t.Errorf("Expected to find CompA in map")
	}
	if compA.Name != "CompA" {
		t.Errorf("Expected name of first component to be %v, but got %v", "CompA", compA.Name)
	}
	if compA.Version != "0.4.2" {
		t.Errorf("Expected version of first component to be %v, but got %v", "0.4.2", compA.Version)
	}
	if compA.License != "MIT, Apache-2.0" {
		t.Errorf("Expected license of first component to be %v, but got %v", "MIT, Apache-2.0", compA.License)
	}

	_, ok = comps[":CompB:1.1.1"]
	if !ok {
		t.Errorf("Expected to find CompB in map")
	}
}

func TestCompMapToSlice(t *testing.T) {
	a := model.Component{
		Name:    "A",
		Pkg:     "a",
		Version: "alpha",
	}
	b := model.Component{
		Name:    "B",
		Pkg:     "b",
		Version: "beta",
	}

	m := mapComp{
		a.ID(): a,
		b.ID(): b,
	}

	comps := compMapToSlice(m)

	if len(comps) != 2 {
		t.Errorf("Expected slice length to be %v, but got %v", 2, len(comps))
	}
	if !model.ContainsComp(comps, a.ID()) {
		t.Errorf("Expected component a to be present in slice")
	}
	if !model.ContainsComp(comps, b.ID()) {
		t.Errorf("Expected component b to be present in slice")
	}
}
