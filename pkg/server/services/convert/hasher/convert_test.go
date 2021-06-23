// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package hasher

import (
	"bytes"
	convert2 "github.com/osrgroup/product-model-toolkit/pkg/services/importing/convert"
	"os"
	"testing"

	"github.com/osrgroup/product-model-toolkit/model"
)

const testFile = "test/example.json"

var artSRC model.Artifact = model.Artifact{
	Path:  "/src",
	Name:  "src",
	IsDir: true,
}

var artClient model.Artifact = model.Artifact{
	Path:  "/src/client",
	Name:  "client",
	IsDir: true,
}

var artMain model.Artifact = model.Artifact{
	Path:  "/src/cli.go",
	Name:  "cli.go",
	IsDir: false,
	Hash: model.Hash{
		MD5:    "9c433b223840fe0aa7977ca9b7bcde7a",
		SHA1:   "3d28572722db70e9d9e650df84a58c69d4b2bead",
		SHA256: "1b12c9e34301701b471ae829d4c27411f7c88dbaf90cf343418da2452e0d4a74",
	},
}

var artCrawler model.Artifact = model.Artifact{
	Path:  "/src/client/crawler.go",
	Name:  "crawler.go",
	IsDir: false,
	Hash: model.Hash{
		MD5:    "9667c36b8eb538d78445614380a8c933",
		SHA1:   "bcd0438536b354f030b259dc8522c9e42903db6d",
		SHA256: "f2cbdaedc0bbafc56dbc006ee537fb27086929424ec6770576366bc3a45ac379",
	},
}

func TestConvert(t *testing.T) {
	jsonFile, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Unable to read %s to start tests", testFile)
	}
	defer jsonFile.Close()

	var c convert2.Converter = new(Hasher)
	p, err := c.Convert(jsonFile)

	t.Run("not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Expected convert func to not return an error, but got %v", err.Error())
		}
	})

	t.Run("not nil", func(t *testing.T) {
		if p == nil {
			t.Error("Expected product to be not nil")
		}
	})

	t.Run("product name", func(t *testing.T) {
		if p.Name != "new Product" {
			t.Errorf("Expected amount of components to be %v, but got %v", "new Product", p.Name)
		}
	})

	t.Run("component amount", func(t *testing.T) {
		if len(p.Components) != 18 {
			t.Errorf("Expected amount of components to be %v, but got %v", 18, len(p.Components))
		}
	})

	t.Run("contains base path as first component", func(t *testing.T) {
		first := p.Components[0]
		if first.Name != "input" {
			t.Errorf("Expected first component name to be '%v', but got '%v'", "input", first.Name)
		}
	})

	t.Run("contains store_test.go as last component", func(t *testing.T) {
		last := p.Components[len(p.Components)-1]
		if last.Name != "store_test.go" {
			t.Errorf("Expected last component name to be '%v', but got '%v'", "store_test.go", last.Name)
		}
	})
}

func TestConvert_Empty(t *testing.T) {
	h := &Hasher{}
	p, err := h.Convert(bytes.NewReader(nil))
	if err == nil {
		t.Error("Expected returning an error for empty doc")
	}
	if p != nil {
		t.Error("Expected resulting product for empty doc to be nil")
	}
}

func TestAsProductModel(t *testing.T) {
	artifacts := []model.Artifact{artSRC, artMain, artCrawler}

	prod, err := asProductModel(artifacts)
	if err != nil {
		t.Errorf("Expected err to be nil, but got %s", err.Error())
	}

	if prod.Name != "new Product" {
		t.Errorf("Expected product name to be 'new Product', but go %v", prod.Name)
	}

	if len(prod.Components) != 3 {
		t.Errorf("Expected amount of component to be 3, but got %v", len(prod.Components))
	}
}

func TestAsProductModel_emptyArtifacts(t *testing.T) {
	_, err := asProductModel([]model.Artifact{})
	if err == nil {
		t.Error("Expected an error with a empty artifact array as argument")
	}
}

func TestAsComponent_withDir(t *testing.T) {
	comp := asComponent(artSRC)

	expectedName := "src"
	if comp.Name != expectedName {
		t.Errorf("Expected component name to be '%v', but got '%v'", expectedName, comp.Name)
	}

	expectedPath := "/src"
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

	expectedPath := "/src/client"
	if comp.Pkg != expectedPath {
		t.Errorf("Expected component pkg to be '%v', but got '%v'", expectedName, comp.Pkg)
	}
}

func TestAsComponent_withFile(t *testing.T) {
	comp := asComponent(artMain)

	expectedName := "cli.go"
	if comp.Name != expectedName {
		t.Errorf("Expected component name to be '%v', but got '%v'", expectedName, comp.Name)
	}

	expectedPath := "/src/cli.go"
	if comp.Pkg != expectedPath {
		t.Errorf("Expected component pkg to be '%v', but got '%v'", expectedName, comp.Pkg)
	}

	if comp.Artifact.Hash.MD5 != artMain.Hash.MD5 {
		t.Errorf("Expected artifact MD5 hash to be '%v', but got '%v'", artMain.Hash.MD5, comp.Artifact.Hash.MD5)
	}

	if comp.Artifact.Hash.SHA1 != artMain.Hash.SHA1 {
		t.Errorf("Expected artifact SHA1 hash to be '%v', but got '%v'", artMain.Hash.SHA1, comp.Artifact.Hash.SHA1)
	}

	if comp.Artifact.Hash.SHA256 != artMain.Hash.SHA256 {
		t.Errorf("Expected artifact MD5 hash to be '%v', but got '%v'", artMain.Hash.SHA256, comp.Artifact.Hash.SHA256)
	}
}

func TestAsComponent_withFileNested(t *testing.T) {
	comp := asComponent(artCrawler)

	expectedName := "crawler.go"
	if comp.Name != expectedName {
		t.Errorf("Expected component name to be '%v', but got '%v'", expectedName, comp.Name)
	}

	expectedPath := "/src/client/crawler.go"
	if comp.Pkg != expectedPath {
		t.Errorf("Expected component pkg to be '%v', but got '%v'", expectedName, comp.Pkg)
	}

	if comp.Artifact.Hash.MD5 != artCrawler.Hash.MD5 {
		t.Errorf("Expected artifact MD5 hash to be '%v', but got '%v'", artCrawler.Hash.MD5, comp.Artifact.Hash.MD5)
	}

	if comp.Artifact.Hash.SHA1 != artCrawler.Hash.SHA1 {
		t.Errorf("Expected artifact SHA1 hash to be '%v', but got '%v'", artCrawler.Hash.SHA1, comp.Artifact.Hash.SHA1)
	}

	if comp.Artifact.Hash.SHA256 != artCrawler.Hash.SHA256 {
		t.Errorf("Expected artifact MD5 hash to be '%v', but got '%v'", artCrawler.Hash.SHA256, comp.Artifact.Hash.SHA256)
	}
}

func TestRemoveBasePath(t *testing.T) {
	type args struct {
		basePath string
		path     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 Level",
			args: args{basePath: "/myProj", path: "/myProj/src/cli.go"},
			want: "/src/cli.go",
		},
		{
			name: "2 Level",
			args: args{basePath: "/myProj", path: "/myProj/src/client/crawler.go"},
			want: "/src/client/crawler.go",
		},
		{
			name: "Root path file",
			args: args{basePath: "/", path: "/cli.go"},
			want: "/cli.go",
		},
		{
			name: "Root path folder with file",
			args: args{basePath: "/", path: "/src/cli.go"},
			want: "/src/cli.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBasePath(tt.args.path, tt.args.basePath); got != tt.want {
				t.Errorf("removeBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
