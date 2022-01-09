// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"github.com/osrgroup/product-model-toolkit/cnst"
	"os"
	"log"
	"github.com/jinzhu/gorm"

	"github.com/osrgroup/product-model-toolkit/pkg/http/rest"
)

var gitCommit string


// Migrate database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Component{})
	db.AutoMigrate(&model.DepGraph{})
	db.AutoMigrate(&model.UsageType{})
	db.AutoMigrate(&model.License{})
}

// Main function
func main() {

	// Check flags
	if checkFlags() {
		os.Exit(0)
	}

    // Connect to database
	db, err := model.Init()
	if err != nil {
		log.Fatalf("db connection err: %v", err)
	}
	defer db.Close()
	
	// Migrate database
	Migrate(db)
	
	// Create services
	repo := model.NewRepo()
	
	// Create REST API port
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = cnst.DefaultServerPort
	}
	
	// Create REST API
	r := rest.NewSrv(
		fmt.Sprintf(":%v", serverPort),
		services.NewService(repo),
	)
	go r.Start()
	defer r.Shutdown()
}

// Check flags
func checkFlags() bool {
	// Print version
	version := flag.Bool("v", false, "show version")

	// Parse flags
	flag.Parse()

	// Print version
	if *version {
		printVersion()
	}

	// Return
	return flag.NFlag() > 0
}

// Print version
func printVersion() {
	fmt.Println("PMT Server")
	fmt.Println("----------")
	fmt.Println("Version: " + "1.0.0")
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}
