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
	version    *bool
	lstScanner *bool
	scanner    string
	inDir      string
	set        bool
}

func main() {
	flg := initFlags()
	did := printInfos(&flg)
	if did {
		os.Exit(0)
	}

	scn := scanner.FromStr(flg.scanner)
	cfg := &scanner.Config{Tool: scn, InDir: flg.inDir, ResultDir: "/tmp/pm/"}

	scanning.Run(cfg)

}

func initFlags() flags {
	version := flag.Bool("v", false, "show version")
	lstScanner := flag.Bool("l", false, "list all available scanner")
	scanner := flag.String("s", "", "scanner to to use from list of available scanner")
	wd, _ := os.Getwd()
	inDir := flag.String("i", wd, "input dir to scan. Default is current working directory")

	flag.Parse()

	set := flag.NFlag() > 0

	return flags{
		version,
		lstScanner,
		*scanner,
		*inDir,
		set,
	}
}

func printInfos(flg *flags) bool {
	if *flg.version {
		printVersion()
		return true
	}

	if *flg.lstScanner {
		listScanner()
		return true
	}

	return false
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
