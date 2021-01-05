// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0
package plugin

import (
	"strings"
	"testing"
)

const registryFileStr = `
[
  {
    "Name": "Licensee",
    "Version": "9.13.0",
    "DockerImg": "docker.pkg.github.com/cmgl/product-model-toolkit/scanner-licensee:9.13.0",
    "Cmd": "/bin/bash -c licensee detect /input/ --json > /result/result.json",
    "Results": [
      "result.json"
    ]
  },
  {
    "Name": "Scancode",
    "Version": "3.1.1",
    "DockerImg": "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1",
    "Cmd": "/bin/bash -c ./scancode --spdx-tv /result/result.spdx /input",
    "Results": [
      "result.spdx"
    ]
  },
  {
    "Name": "Composer",
    "Version": "dummy",
    "DockerImg": "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy",
    "Cmd": "/bin/sh -c cp example.json result/example.json",
    "Results": [
      "example.json"
    ]
  }
]
`

func TestIsEmpty(t *testing.T) {
	emptyReg := Registry{}
	isEmpty := emptyReg.IsEmpty()
	if isEmpty == false {
		t.Error("Expected isEmpty() to return true for empty registry")
	}

	populatedReg := generateRegistry()
	isEmpty = populatedReg.IsEmpty()
	if isEmpty == true {
		t.Error("Expected isEmpty() for populated registry to return false")
	}
}

func TestDoImportFromFile(t *testing.T) {
	handler := strings.NewReader(registryFileStr)
	plugins, err := doImportFromFile(handler)
	if err != nil {
		t.Errorf("Expected import from file without err, but got %s", err.Error())
	}

	if len(plugins) != 3 {
		t.Errorf("Expected plugins length of imported registry to be 3, but got %v", len(plugins))
	}
}

func TestDoImportFromFile_EmptyFile(t *testing.T) {
	handler := strings.NewReader("")
	plugins, err := doImportFromFile(handler)
	if err == nil {
		t.Error("Expected import ro return error for empty string input")
	}

	if len(plugins) != 0 {
		t.Errorf("Expected plugins length of empty registry to be 0, but got %v", len(plugins))
	}
}

func TestFromStr(t *testing.T) {
	var reg Register = generateRegistry()

	var name string
	name = "Licensee"
	plugin, found := reg.FromStr(name)
	checkFoundPlugin(found, name, t)
	checkPluginName(name, plugin.Name, t)

	name = "Scancode"
	plugin, found = reg.FromStr(name)
	checkFoundPlugin(found, name, t)
	checkPluginName(name, plugin.Name, t)

	name = "Composer"
	plugin, found = reg.FromStr(name)
	checkFoundPlugin(found, name, t)
	checkPluginName(name, plugin.Name, t)
}

func TestFromStr_CaseInsensitive(t *testing.T) {
	var reg Register = generateRegistry()
	licensee, found := reg.FromStr("lIceNsEE")
	scancode, found := reg.FromStr("scANcoDE")

	if found == false || licensee.Name != "Licensee" {
		t.Error("Expected FromStr() input to work case insensitive")
	}

	if found == false || scancode.Name != "Scancode" {
		t.Error("Expected FromStr() input to work case insensitive")
	}
}

func TestFromStr_EmptyRegistry(t *testing.T) {
	var reg Register = &Registry{plugins: []Plugin{}}

	_, found := reg.FromStr("Licensee")
	if found == true {
		t.Error("Expected FromStr() for an empty registry to not return false")
	}

	_, found = reg.FromStr("")
	if found == true {
		t.Error("Expected FromStr() for an empty registry to not return false")
	}
}

func TestAvailable(t *testing.T) {
	reg := generateRegistry()
	plugins := reg.Available()
	if len(plugins) != 3 {
		t.Errorf("Expected length of available plugins to be 3, but got %v", len(plugins))
	}
}

func TestDefault(t *testing.T) {
	reg := generateRegistry()

	first := reg.Default()
	expected := "Licensee"
	if first.Name != expected {
		t.Errorf("Expected name of default plugin to be '%s', but got '%s'", expected, first.Name)
	}
}

func TestDefault_EmptyRegistry(t *testing.T) {
	reg := &Registry{}

	first := reg.Default()
	if first.Name != "" {
		t.Errorf("Expected to return empty plugin if registry is empty, but got plugin with name '%s'", first.Name)
	}
}

func checkFoundPlugin(found bool, name string, t *testing.T) {
	if found == false {
		t.Errorf("Expected FromStr() to find plugin with name '%s'", name)
	}
}

func checkPluginName(expected, name string, t *testing.T) {
	if expected != name {
		t.Errorf("Expected FromStr() find plugin with name '%s', but got '%s'", expected, name)
	}
}

func generateRegistry() *Registry {
	plugins, _ := doImportFromFile(strings.NewReader(registryFileStr))
	return &Registry{
		plugins: plugins,
	}
}
