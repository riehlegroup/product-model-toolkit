// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

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

// init loads available scanner tools from config file. It also assigns default tool
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "scanner_tools.json")
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

// NoTools returns true if no scanner tools available
func NoTools() bool {
	return len(Available) <= 0
}

// FromStr returns a tool with the given name and indicates with a bool if the tool could be found
func FromStr(name string) (Tool, bool) {
	search := strings.ToLower(name)
	for _, t := range Available {
		if strings.ToLower(t.Name) == search {
			return t, true
		}
	}

	return Tool{}, false
}

// String return the name and the version of the tool.
func (t *Tool) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Version)
}
