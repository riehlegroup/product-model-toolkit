// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	// "github.com/osrgroup/product-model-toolkit/pkg/db/memory"
	"github.com/osrgroup/product-model-toolkit/model"
	// "github.com/osrgroup/product-model-toolkit/pkg/db/postgraph"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"github.com/osrgroup/product-model-toolkit/cnst"
	"os"
	"log"
	"github.com/jinzhu/gorm"

	"github.com/osrgroup/product-model-toolkit/pkg/http/rest"
)

var gitCommit string


func Migrate(db *gorm.DB) {
	// users.AutoMigrate()
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Component{})
	db.AutoMigrate(&model.DepGraph{})
	db.AutoMigrate(&model.UsageType{})
	db.AutoMigrate(&model.License{})

	// db.AutoMigrate(&articles.TagModel{})
	// db.AutoMigrate(&articles.FavoriteModel{})
	// db.AutoMigrate(&articles.ArticleUserModel{})
	// db.AutoMigrate(&articles.CommentModel{})
}

func main() {
	if checkFlags() {
		os.Exit(0)
	}

	db, err := model.Init()
	if err != nil {
		log.Fatalf("db connection err: %v", err)
	}
	Migrate(db)
	defer db.Close()
	// repo := new(memory.DB)
	// repo.AddSampleData()

	// Use Postgraphile as DB backend
	// repo := postgraph.NewRepo("http://127.0.0.1:5433/graphql")
	repo := model.NewRepo()
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = cnst.DefaultServerPort
	}
	
	r := rest.NewSrv(
		fmt.Sprintf(":%v", serverPort),
		services.NewService(repo),
	)
	go r.Start()
	defer r.Shutdown()
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
	fmt.Println("Version: " + "1.0.0")
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}
