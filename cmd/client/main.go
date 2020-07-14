// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/client/scanner"
	"github.com/osrgroup/product-model-toolkit/pkg/client/scanning"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

type flags struct {
	scanner string
	inDir   string
}

func main() {
	flg, abort := checkFlags()
	if abort {
		os.Exit(0)
	}

	scn := scanner.FromStr(flg.scanner)
	cfg := &scanner.Config{Tool: scn, InDir: flg.inDir, ResultDir: "/tmp/pm/"}

	scanning.Run(cfg)

}

func checkFlags() (flags, bool) {
	version := flag.Bool("v", false, "show version")

	lstScanner := flag.Bool("l", false, "list all available scanner")

	scanner := flag.String("s", "", "scanner to to use from list of available scanner")
	wd, _ := os.Getwd()
	inDir := flag.String("i", wd, "input dir to scan. Default is current working directory")

	flag.Parse()

	if *version {
		printVersion()
	}

	if *lstScanner {
		listScanner()
	}

	abortAfterFlags := *version || *lstScanner

	return flags{
			*scanner,
			*inDir,
		},
		abortAfterFlags
}

func printVersion() {
	fmt.Println("PMT Client")
	fmt.Println("----------")
	fmt.Println("Version: " + version.Name())
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}

func listScanner() {
	fmt.Println("Available license scanner:")
	for _, scn := range scanner.Available {
		fmt.Printf("%+v\n", scn)
	}
}
