// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package composer

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/osrgroup/product-model-toolkit/pkg/convert"
)

var exampleDoc []byte

// TestMain runs before all tests once
func TestMain(m *testing.M) {
	var err error
	exampleDoc, err = readExampleDoc()
	if err != nil {
		log.Fatal("Unable to read example.json to start tests")
		os.Exit(-1)
	}
	os.Exit(m.Run())
}

func readExampleDoc() ([]byte, error) {
	jsonFile, err := os.Open("test/example.json")
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
	t.Run("component name", func(t *testing.T) {
		first := p.Components[0].Name
		firstExp := "bluespice/about"
		last := p.Components[len(p.Components)-1].Name
		lastExp := "zordius/lightncandy"

		if first != firstExp {
			t.Errorf("Expected name of first component to be %v, but got %v", firstExp, first)
		}
		if last != lastExp {
			t.Errorf("Expected name of last component to be %v, but got %v", lastExp, last)
		}
	})

	t.Run("component version", func(t *testing.T) {
		first := p.Components[0].Version
		firstExp := "dev-REL1_31"
		last := p.Components[len(p.Components)-1].Version
		lastExp := "v0.23"

		if first != firstExp {
			t.Errorf("Expected version of first component to be %v, but got %v", firstExp, first)
		}
		if last != lastExp {
			t.Errorf("Expected version of last component to be %v, but got %v", lastExp, last)
		}
	})

	t.Run("component license", func(t *testing.T) {
		first := p.Components[0].License
		firstExp := "GPL-3.0-only"
		last := p.Components[len(p.Components)-1].License
		lastExp := "MIT"

		if first != firstExp {
			t.Errorf("Expected license of first component to be %v, but got %v", firstExp, first)
		}
		if last != lastExp {
			t.Errorf("Expected license of last component to be %v, but got %v", lastExp, last)
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

	comps := extractComponents(input)
	if comps == nil {
		t.Errorf("Expected extractComponents() result not be nil")
	}
	if len(comps) != 2 {
		t.Errorf("Expected to return %v components, but got %v", 2, len(comps))
	}
	if comps[0].Name != "CompA" {
		t.Errorf("Expected name of first component to be %v, but got %v", "CompA", comps[0].Name)
	}
	if comps[0].Version != "0.4.2" {
		t.Errorf("Expected version of first component to be %v, but got %v", "0.4.2", comps[0].Version)
	}
	if comps[0].License != "MIT, Apache-2.0" {
		t.Errorf("Expected license of first component to be %v, but got %v", "MIT, Apache-2.0", comps[0].License)
	}
	if comps[1].License != "GPL" {
		t.Errorf("Expected license of second component to be %v, but got %v", "GPL", comps[1].License)
	}

}
