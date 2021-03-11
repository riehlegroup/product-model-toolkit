// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/client/plugin"
	"github.com/osrgroup/product-model-toolkit/pkg/client/scanning"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

const serverBaseURL = "http://localhost:8081/api/v1"

type flags struct {
	scanner string
	regFile string
	inDir   string
}

func main() {
	flg, abort := checkFlags()
	if abort {
		os.Exit(0)
	}

	var pluginRegistry plugin.Register = loadPluginRegistry(flg.regFile)

	scn, found := pluginRegistry.FromStr(flg.scanner)
	if !found {
		scn = pluginRegistry.Default()
		log.Printf("[Core] Unable to find scanner plugin with name '%s'; fallback to default scanner with name '%s'", flg.scanner, scn.Name)
	}

	tempDir, err := ioutil.TempDir("", "pm-*")
	if err != nil {
		log.Print("[Core] Unable to create a temporary directory\nUnable to proceed")
		os.Exit(-1)
	}

	cfg := &plugin.Config{Plugin: scn, InDir: flg.inDir, ResultDir: tempDir}

	scanning.Run(
		cfg,
		rest.NewClient(serverBaseURL),
	)
}

func checkFlags() (flags, bool) {
	version := flag.Bool("v", false, "show version")
	lstScanner := flag.Bool("l", false, "list all available scanner plugins")
	regFile := flag.String("r", "plugins.yml", "plugin registry file to use")
	scanner := flag.String("s", "", "scanner to to use from list of available scanner")

	wd, _ := os.Getwd()
	inDir := flag.String("i", wd, "input dir to scan. Default is current working directory")

	flag.Parse()

	if *version {
		printVersion()
	}

	if *lstScanner {
		listScanner(*regFile)
	}

	abortAfterFlags := *version || *lstScanner

	return flags{
			*scanner,
			*regFile,
			*inDir,
		},
		abortAfterFlags
}

func printVersion() {
	fmt.Printf(
		"PMT Client\n----------\nVersion: %s\nGit commit: %s\n",
		version.Name(),
		gitCommit,
	)
}

func loadPluginRegistry(file string) *plugin.Registry {
	pluginRegistry, err := plugin.NewRegistry(file)
	if err != nil {
		log.Printf("[Core] Unable to create new plugin registry from file '%s'. Error: %s\nUnable to proceed", file, err.Error())
		os.Exit(-1)
	}
	if pluginRegistry.IsEmpty() {
		log.Print("[Core] Unable to proceed with empty plugin registry")
		os.Exit(-1)
	}
	return pluginRegistry
}

func listScanner(regFile string) {
	pluginRegistry := loadPluginRegistry(regFile)
	plugins := pluginRegistry.Available()
	fmt.Printf("Available license scanner from plugin file '%s':\n", regFile)
	for _, scn := range plugins {
		fmt.Printf("----------\nName:    %s\nVersion: %s\nImage:   %s\n", scn.Name, scn.Version, scn.DockerImg)
	}
	fmt.Printf("----------\n")
}
