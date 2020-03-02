package main

import (
	"flag"
	"fmt"

	"github.com/osrgroup/product-model-toolkit/pkg/scanner"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

func main() {
	initFlags()

	fmt.Println("Hello, I'm the client.")
}

func initFlags() {
	version := flag.Bool("v", false, "show version")
	list := flag.Bool("l", false, "list all available scanner")

	flag.Parse()

	if *version {
		printVersion()
	}

	if *list {
		listScanner()
	}
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
