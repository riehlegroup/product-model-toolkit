// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0
package plugin

import (
	"strings"
	"testing"
)

const registryFileStrYaml = `
version: R1.0
defaultplugin: 0
plugins:
  - name: Licensee
    version: 9.13.0
    dockerimg: docker.pkg.github.com/cmgl/product-model-toolkit/scanner-licensee:9.13.0
    shell: /bin/bash
    cmd: licensee detect /input/ --json > /result/result.json
    results:
      - result.json
  - name: Scancode
    version: 3.1.1
    dockerimg: docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1
    shell: /bin/bash
    cmd: ./scancode --spdx-tv /result/result.spdx /input
    results:
      - result.spdx
  - name: Composer
    version: dummy
    dockerimg: docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy
    shell: /bin/sh
    cmd: cp example.json result/example.json
    results:
      - example.json
`

const registryFileStrJson = `
{
  "Version": "R1.0",
  "DefaultPlugin": 0,
  "Plugins": [
    {
      "Name": "Licensee",
      "Version": "9.13.0",
      "DockerImg": "docker.pkg.github.com/cmgl/product-model-toolkit/scanner-licensee:9.13.0",
      "Shell": "/bin/bash",
      "Cmd": "licensee detect /input/ --json \u003e /result/result.json",
      "Results": [
        "result.json"
      ]
    },
    {
      "Name": "Scancode",
      "Version": "3.1.1",
      "DockerImg": "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1",
      "Shell": "/bin/bash",
      "Cmd": "./scancode --spdx-tv /result/result.spdx /input",
      "Results": [
        "result.spdx"
      ]
    },
    {
      "Name": "Composer",
      "Version": "dummy",
      "DockerImg": "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy",
      "Shell": "/bin/sh",
      "Cmd": "cp example.json result/example.json",
      "Results": [
        "example.json"
      ]
    }
  ]
}
`

func TestDoImportFromYamlFile(t *testing.T) {
	handler := []byte(registryFileStrYaml)
	registry, err := doImportFromYamlFile(handler)
	if err != nil {
		t.Errorf("Expected import from file without err, but got %s", err.Error())
	}

	if len(registry.Plugins) != 3 {
		t.Errorf("Expected plugins length of imported registry to be 3, but got %v", len(registry.Plugins))
	}
}

func TestDoImportFromYamlFile_EmptyFile(t *testing.T) {
	handler := []byte("")
	registry, err := doImportFromYamlFile(handler)
	if err == nil {
		t.Error("Expected import ro return error for empty string input")
	}

	if len(registry.Plugins) != 0 {
		t.Errorf("Expected plugins length of empty registry to be 0, but got %v", len(registry.Plugins))
	}
}

func TestDoImportFromJsonFile(t *testing.T) {
	handler := strings.NewReader(registryFileStrJson)
	registry, err := doImportFromJsonFile(handler)
	if err != nil {
		t.Errorf("Expected import from file without err, but got %s", err.Error())
	}

	if len(registry.Plugins) != 3 {
		t.Errorf("Expected plugins length of imported registry to be 3, but got %v", len(registry.Plugins))
	}
}

func TestDoImportFromJsonFile_EmptyFile(t *testing.T) {
	handler := strings.NewReader("")
	registry, err := doImportFromJsonFile(handler)
	if err == nil {
		t.Error("Expected import ro return error for empty string input")
	}

	if len(registry.Plugins) != 0 {
		t.Errorf("Expected plugins length of empty registry to be 0, but got %v", len(registry.Plugins))
	}
}

func TestAvailable(t *testing.T) {
	reg := generateRegistryFromYamlFile()
	plugins := reg.Available()
	if len(plugins) != 3 {
		t.Errorf("Expected length of available plugins to be 3, but got %v", len(plugins))
	}
}

func TestDefault(t *testing.T) {
	reg := generateRegistryFromYamlFile()

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

func TestIsEmpty(t *testing.T) {
	emptyReg := Registry{}
	isEmpty := emptyReg.IsEmpty()
	if isEmpty == false {
		t.Error("Expected isEmpty() to return true for empty registry")
	}

	populatedReg := generateRegistryFromYamlFile()
	isEmpty = populatedReg.IsEmpty()
	if isEmpty == true {
		t.Error("Expected isEmpty() for populated registry (from yaml file) to return false")
	}

	populatedReg = generateRegistryFromJsonFile()
	isEmpty = populatedReg.IsEmpty()
	if isEmpty == true {
		t.Error("Expected isEmpty() for populated registry (from json file) to return false")
	}
}

func TestFromStr(t *testing.T) {
	var reg Register = generateRegistryFromYamlFile()

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
	var reg Register = generateRegistryFromYamlFile()
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
	var reg Register = &Registry{Plugins: []Plugin{}}

	_, found := reg.FromStr("Licensee")
	if found == true {
		t.Error("Expected FromStr() for an empty registry to not return false")
	}

	_, found = reg.FromStr("")
	if found == true {
		t.Error("Expected FromStr() for an empty registry to not return false")
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

func generateRegistryFromYamlFile() *Registry {
	registry, _ := doImportFromYamlFile([]byte(registryFileStrYaml))
	return &Registry{
		Plugins: registry.Plugins,
	}
}

func generateRegistryFromJsonFile() *Registry {
	registry, _ := doImportFromJsonFile(strings.NewReader(registryFileStrJson))
	return &Registry{
		Plugins: registry.Plugins,
	}
}
