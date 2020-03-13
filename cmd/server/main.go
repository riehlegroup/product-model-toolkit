package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/server"
	"github.com/osrgroup/product-model-toolkit/pkg/server/dgraph"
	"github.com/osrgroup/product-model-toolkit/pkg/server/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

func main() {
	usedFlag := initFlags()

	if usedFlag {
		os.Exit(0)
	}

	r := rest.NewSrv("127.0.0.1:8080")
	go r.Start()
	defer r.Shutdown()

	db := dgraph.NewClient(dgraph.DefaultURI)

	srv := &server.Instance{
		REST: r,
		DB:   db,
	}

	srv.DB.DropAll()
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
