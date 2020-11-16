// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

// Plugin struct defines a (scanner) plugin
type Plugin struct {
	Name      string
	Version   string
	DockerImg string
	Cmd       string
	Results   []string
}

// Available is a list of available plugins that can be used
var Available []Plugin

// Default is the default (scanner) plugin that shall be used if no particular plugin is selected
var Default Plugin

// init loads available (scanner) plugins from plugin registry and assigns default plugin
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "plugin_registry.json")
	jsonFile, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	err = json.Unmarshal(byteValue, &Available)
	if err != nil {
		log.Println(err)
	}

	Default = Available[0]
}

// NoPlugins returns true if no (scanner) plugins available
func NoPlugins() bool {
	return len(Available) <= 0
}

// FromStr returns a plugin with the given name and indicates with a bool if plugin is found
func FromStr(name string) (Plugin, bool) {
	search := strings.ToLower(name)
	for _, p := range Available {
		if strings.ToLower(p.Name) == search {
			return p, true
		}
	}

	return Plugin{}, false
}

// String returns the name and the version of the plugin
func (p *Plugin) String() string {
	return fmt.Sprintf("%s (%s)", p.Name, p.Version)
}
