// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
)

// Plugin struct defines a plugin
type Plugin struct {
	Name      string
	Version   string
	DockerImg string
	Shell     string
	Cmd       string
	Results   []string
}

// Config represents a configuration for a plugin to execute
type Config struct {
	Plugin
	InDir     string
	ResultDir string
	Id        int
}

// String returns the name and the version of the plugin
func (p *Plugin) String() string {
	return fmt.Sprintf("%s (%s)", p.Name, p.Version)
}
