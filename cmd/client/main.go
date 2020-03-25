package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/scanner"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

func main() {
	usedFlag := initFlags()
	if usedFlag {
		os.Exit(0)
	}

	fmt.Println("Hello, I'm the client.")
}

func initFlags() bool {
	version := flag.Bool("v", false, "show version")
	list := flag.Bool("l", false, "list all available scanner")

	flag.Parse()

	if *version {
		printVersion()
	}

	if *list {
		listScanner()
	}

	return flag.NFlag() > 0
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
