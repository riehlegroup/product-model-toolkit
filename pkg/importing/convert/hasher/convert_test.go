// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package hasher

import "testing"

const testFile = "test/example.json"

var artSRC artifact = artifact{
	Path:  "src",
	Name:  "src",
	IsDir: true,
}

var artClient artifact = artifact{
	Path:  "src/client",
	Name:  "client",
	IsDir: true,
}

var artMain artifact = artifact{
	Path:  "src/main.go",
	Name:  "main.go",
	IsDir: false,
	Hash: hash{
		MD5:    "9c433b223840fe0aa7977ca9b7bcde7a",
		SHA1:   "3d28572722db70e9d9e650df84a58c69d4b2bead",
		SHA256: "1b12c9e34301701b471ae829d4c27411f7c88dbaf90cf343418da2452e0d4a74",
	},
}

var artCrawler artifact = artifact{
	Path:  "src/client/crawler.go",
	Name:  "crawler.go",
	IsDir: false,
	Hash: hash{
		MD5:    "9667c36b8eb538d78445614380a8c933",
		SHA1:   "bcd0438536b354f030b259dc8522c9e42903db6d",
		SHA256: "f2cbdaedc0bbafc56dbc006ee537fb27086929424ec6770576366bc3a45ac379",
	},
}

func TestAsProductModel(t *testing.T) {
	artifacts := []artifact{artSRC, artMain, artCrawler}

	prod, err := asProductModel(artifacts)
	if err != nil {
		t.Errorf("Expected err to be nil, but got %s", err.Error())
	}

	if prod.Name != "new Product" {
		t.Errorf("Expected product name to be 'new Product', but go %v", prod.Name)
	}

	if len(prod.Components) != 2 {
		t.Errorf("Expected amount of component to be 2, but got %v", len(prod.Components))
	}
}

func TestAsComponent_withDir(t *testing.T) {
	comp := asComponent(artSRC)

	expectedName := "src"
	if comp.Name != expectedName {
		t.Errorf("Expected component name to be '%v', but got '%v'", expectedName, comp.Name)
	}

	expectedPath := "src"
	if comp.Pkg != expectedPath {
		t.Errorf("Expected component pkg to be '%v', but got '%v'", expectedName, comp.Pkg)
	}
}

func TestAsComponent_withDirNested(t *testing.T) {
	comp := asComponent(artClient)

	expectedName := "client"
	if comp.Name != expectedName {
		t.Errorf("Expected component name to be '%v', but got '%v'", expectedName, comp.Name)
	}

	expectedPath := "src/client"
	if comp.Pkg != expectedPath {
		t.Errorf("Expected component pkg to be '%v', but got '%v'", expectedName, comp.Pkg)
	}
}
