// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

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
	plugins, err := importFromFile(file)
	if err != nil {
		return &Registry{}, err
	}

	return &Registry{plugins: plugins}, nil
}

// importFromFile parses a given JSON registry file into []Plugin
func importFromFile(file string) ([]Plugin, error) {
	handle, err := os.Open(file)
	if err != nil {
		return []Plugin{}, err
	}
	defer handle.Close()

	return doImportFromFile(handle)
}

// doImportFromFile parses a given JSON registry io stream into []Plugin
func doImportFromFile(handler io.Reader) ([]Plugin, error) {
	var plugins []Plugin
	err := json.NewDecoder(handler).Decode(&plugins)
	if err != nil {
		return []Plugin{}, err
	}

	return plugins, nil
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
