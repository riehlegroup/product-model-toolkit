package scanner

import (
	"fmt"
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
var Available = [...]Tool{
	Licensee,
	Scancode,
}

// Default is the default scanner tools that shall be used if no particular tool is selected.
var Default Tool = Licensee

// Licensee represents the latest usable Licensee version
var Licensee = Tool{
	Name:      "Licensee",
	Version:   "9.13.0",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0",
	Cmd:       "licensee detect /input/hamster/ --json > /results/scan.json",
	Results:   []string{"scan.json"},
}

// Scancode represents the latest usable Scancode version
var Scancode = Tool{
	Name:      "Scancode",
	Version:   "3.1.1",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1",
	Cmd:       "scancode -v",
	Results:   []string{"result.spdx"},
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
