package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/db/memory"
	"github.com/osrgroup/product-model-toolkit/pkg/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

func main() {
	usedFlag := initFlags()

	if usedFlag {
		os.Exit(0)
	}

	repo := new(memory.DB)
	repo.AddSampleData()

	querying := querying.NewService(repo)

	r := rest.NewSrv("127.0.0.1:8081", &querying)
	go r.Start()
	defer r.Shutdown()
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
