// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const registryVersion = "R1.0"

// Registry represents a plugin registry
type Registry struct {
	Version       string
	DefaultPlugin int
	Plugins       []Plugin
}

// Register provides plugin registry operations
type Register interface {
	Available() []Plugin
	Default() Plugin
	IsEmpty() bool
	FromStr(string) (Plugin, bool)
}

// NewRegistry returns a new plugin registry from YAML or JSON input file
func NewRegistry(file string) (*Registry, error) {
	if strings.Contains(file, ".yml") {
		registry, err := importFromYamlFile(file)
		if err != nil {
			return &Registry{}, err
		}
		if registry.Version == registryVersion {
			return &registry, nil
		}
	}
	if strings.Contains(file, ".json") {
		registry, err := importFromJsonFile(file)
		if err != nil {
			return &Registry{}, err
		}
		if registry.Version == registryVersion {
			return &registry, nil
		}
	}

	return &Registry{}, errors.New("unsupported config file format or version")
}

// importFromYamlFile parses a given YAML registry file into []Plugin
func importFromYamlFile(file string) (Registry, error) {
	handle, err := ioutil.ReadFile(file)
	if err != nil {
		return Registry{}, err
	}

	return doImportFromYamlFile(handle)
}

// doImportFromYamlFile parses a given YAML registry byte slice into []Plugin
func doImportFromYamlFile(handler []byte) (Registry, error) {
	if len(handler) == 0 {
		return Registry{}, errors.New("file is empty")
	}
	var registry Registry
	if err := yaml.Unmarshal(handler, &registry); err != nil {
		return Registry{}, err
	}

	return registry, nil
}

// importFromJsonFile parses a given JSON registry file into []Plugin
func importFromJsonFile(file string) (Registry, error) {
	handle, err := os.Open(file)
	if err != nil {
		return Registry{}, err
	}
	defer handle.Close()

	return doImportFromJsonFile(handle)
}

// doImportFromJsonFile parses a given JSON registry io stream into []Plugin
func doImportFromJsonFile(handler io.Reader) (Registry, error) {
	var registry Registry
	err := json.NewDecoder(handler).Decode(&registry)
	if err != nil {
		return Registry{}, err
	}

	return registry, nil
}

// Available returns all plugins in the registry
func (r *Registry) Available() []Plugin {
	return r.Plugins
}

// Default returns the default plugin in the registry
func (r *Registry) Default() Plugin {
	if r.IsEmpty() {
		return Plugin{}
	}

	return r.Plugins[r.DefaultPlugin]
}

// IsEmpty returns true if no plugins are available
func (r *Registry) IsEmpty() bool {
	return len(r.Plugins) <= 0
}

// FromStr returns a plugin with the given name and indicates with a bool if plugin is found
func (r *Registry) FromStr(name string) (Plugin, bool) {
	search := strings.ToLower(name)
	for _, p := range r.Plugins {
		if strings.ToLower(p.Name) == search {
			return p, true
		}
	}

	return Plugin{}, false
}
