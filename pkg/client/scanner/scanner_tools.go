// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

// Composer represents a container which produces a Composer result artifact.
var Composer = Tool{
	Name:      "Composer",
	Version:   "dummy",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy",
	Cmd:       `/bin/sh -c "cp example.json result/example.json"`,
	Results:   []string{"example.json"},
}

// FileHasher represents a container which produces a JSON file with hashes of all file of the input folder.
var FileHasher = Tool{
	Name:      "File-Hasher",
	Version:   "latest",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/file-hasher:latest",
	Cmd:       `/bin/sh -c "./fh -i /input -o /result/result.json"`,
	Results:   []string{"result.json"},
}
