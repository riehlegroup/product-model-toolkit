package main

import (
	"flag"
	"fmt"

	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

func main() {
	initFlags()

	fmt.Println("Hello, I'm the client.")
}

func initFlags() {
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		printVersion()
	}
}

func printVersion() {
	fmt.Println("PMT Client")
	fmt.Println("----------")
	fmt.Println("Version: " + version.Name())
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}
