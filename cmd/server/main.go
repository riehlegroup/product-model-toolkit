package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/version"
	"github.com/osrgroup/product-model-toolkit/pkg/version/server/http/rest"
)

var gitCommit string

func main() {
	usedFlag := initFlags()

	if usedFlag {
		os.Exit(0)
	}

	srv := rest.NewSrv("127.0.0.1:8080")
	go srv.Start()
	defer srv.Shutdown()
}

func initFlags() bool {
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		printVersion()
	}

	return flag.NFlag() > 0
}

func printVersion() {
	fmt.Println("PMT Server")
	fmt.Println("----------")
	fmt.Println("Version: " + version.Name())
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}
