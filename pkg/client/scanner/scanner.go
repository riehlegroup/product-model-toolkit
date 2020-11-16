// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"github.com/osrgroup/product-model-toolkit/pkg/plugin"
)

// Scanner provides license scanner operations.
type Scanner interface {
	Exec(cfg Config)
}

// Config represents a configuration for a tool to execute.
type Config struct {
	plugin.Plugin
	InDir     string
	ResultDir string
}
