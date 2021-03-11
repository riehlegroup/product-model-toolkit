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

// RegistryFile represents a plugin registry including its version number
type RegistryFile struct {
	Version string
	Plugins []Plugin
}

// Registry represents a plugin registry
type Registry struct {
	plugins []Plugin
}

// Register provides plugin registry operations
type Register interface {
	Available() []Plugin
	Default() Plugin
	IsEmpty() bool
	FromStr(string) (Plugin, bool)
}

// NewRegistry returns a new plugin registry from a JSON input file
func NewRegistry(file string) (*Registry, error) {
	if strings.Contains(file, ".yml") {
		plugins, err := importFromYamlFile(file)
		if err != nil {
			return &Registry{}, err
		}
		return &Registry{plugins: plugins}, nil
	}
	if strings.Contains(file, ".json") {
		plugins, err := importFromJsonFile(file)
		if err != nil {
			return &Registry{}, err
		}
		return &Registry{plugins: plugins}, nil
	}

	return &Registry{}, errors.New("unsupported config file format")
}

// importFromYamlFile parses a given YAML registry file into []Plugin
func importFromYamlFile(file string) ([]Plugin, error) {
	handle, err := ioutil.ReadFile(file)
	if err != nil {
		return []Plugin{}, err
	}

	return doImportFromYamlFile(handle)
}

// doImportFromYamlFile parses a given YAML registry io stream into []Plugin
func doImportFromYamlFile(handler []byte) ([]Plugin, error) {
	if len(handler) == 0 {
		return []Plugin{}, errors.New("file is empty")
	}
	var registryFile RegistryFile
	if err := yaml.Unmarshal(handler, &registryFile); err != nil {
		return []Plugin{}, err
	}

	return registryFile.Plugins, nil
}

// importFromJsonFile parses a given JSON registry file into []Plugin
func importFromJsonFile(file string) ([]Plugin, error) {
	handle, err := os.Open(file)
	if err != nil {
		return []Plugin{}, err
	}
	defer handle.Close()

	return doImportFromJsonFile(handle)
}

// doImportFromJsonFile parses a given JSON registry io stream into []Plugin
func doImportFromJsonFile(handler io.Reader) ([]Plugin, error) {
	var registryFile RegistryFile
	err := json.NewDecoder(handler).Decode(&registryFile)
	if err != nil {
		return []Plugin{}, err
	}

	return registryFile.Plugins, nil
}

// IsEmpty returns true if no scanner plugins are available
func (r *Registry) IsEmpty() bool {
	return len(r.plugins) <= 0
}

// FromStr returns a plugin with the given name and indicates with a bool if plugin is found
func (r *Registry) FromStr(name string) (Plugin, bool) {
	search := strings.ToLower(name)
	for _, p := range r.plugins {
		if strings.ToLower(p.Name) == search {
			return p, true
		}
	}

	return Plugin{}, false
}

// Available returns all plugins of the registry
func (r *Registry) Available() []Plugin {
	return r.plugins
}

// Default returns the default plugin of the registry
func (r *Registry) Default() Plugin {
	if r.IsEmpty() {
		return Plugin{}
	}

	return r.plugins[0]
}
