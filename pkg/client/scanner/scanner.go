// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Scanner provides license scanner operations.
type Scanner interface {
	Exec(cfg Config)
}

// Tool struct define a scanner tool
type Tool struct {
	Name      string
	Version   string
	DockerImg string
	Cmd       string
	Results   []string
}

// Config represents a configuration for a tool to execute.
type Config struct {
	Tool
	InDir     string
	ResultDir string
}

// Available list all available scanner tools that can be used.
var Available []Tool

// Default is the default scanner tools that shall be used if no particular tool is selected.
var Default Tool

// Licensee represents the latest usable Licensee version
var Licensee Tool

// Scancode represents the latest usable Scancode version
var Scancode Tool

// init reads data from JSON string to identify available scanner tools
func init() {

	// TODO: Import string from JSON file
	jsonData := `
[
  {
    "Name": "Licensee",
    "Version": "9.13.0",
    "DockerImg": "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0",
    "Cmd": "/bin/bash -c \"licensee detect /input/ --json > /result/result.json\"",
    "Results": [
      "result.json"
    ]
  },
  {
    "Name": "Scancode",
    "Version": "3.1.1",
    "DockerImg": "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1",
    "Cmd": "/bin/bash -c \"./scancode --spdx-tv /result/result.spdx /input\"",
    "Results": [
      "result.spdx"
    ]
  }
]`

	// Decode JSON data and create list of scanner tools
	var Tools []Tool
	err := json.Unmarshal([]byte(jsonData), &Tools)
	if err != nil {
		log.Println(err)
	}

	// Assign scanner tools from the list
	Licensee = Tools[0]
	Scancode = Tools[1]

	// Assign all available scanner tools
	Available = []Tool{
		Licensee,
		Scancode,
		Composer,
		FileHasher,
	}

	// Assign default scanner tool
	Default = Licensee
}

// FromStr returns a tool with the given name. If unable to find a tool with the given name it returns the default tool.
func FromStr(name string) Tool {
	search := strings.ToLower(name)
	for _, t := range Available {
		if strings.ToLower(t.Name) == search {
			return t
		}
	}

	return Default
}

// String return the name and the version of the tool.
func (t *Tool) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Version)
}
