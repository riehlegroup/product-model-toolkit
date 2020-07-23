// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/osrgroup/product-model-toolkit/pkg/db/memory"
	"github.com/osrgroup/product-model-toolkit/pkg/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/importing"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

func main() {
	if checkFlags() {
		os.Exit(0)
	}

	repo := new(memory.DB)
	repo.AddSampleData()
	// Use Postgraphile as DB backend
	// repo := postgraph.NewRepo("http://localhost:5433/graphql")

	r := rest.NewSrv(
		"127.0.0.1:8081",
		querying.NewService(repo),
		importing.NewService(repo),
	)
	go r.Start()
	defer r.Shutdown()

	time.Sleep(300 * time.Millisecond)
	r.PrintRoutes()
}

func checkFlags() bool {
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
