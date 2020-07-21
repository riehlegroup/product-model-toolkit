// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

// ComposerDummy represents a dummy container which contains a Composer result artifact.
var ComposerDummy = Tool{
	Name:      "Composer",
	Version:   "dummy",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy",
	Cmd:       `/bin/sh -c "cp example.json result/example.json"`,
	Results:   []string{"example.json"},
}
