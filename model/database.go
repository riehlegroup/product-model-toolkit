// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() (*gorm.DB, error) {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	conn.DB().SetMaxIdleConns(10)
	conn.LogMode(true)
	DB = conn
	return DB, nil
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

type repo struct {
	conn *gorm.DB
}

// NewRepo create an instance for Postgraphile DB repository
func NewRepo() *repo {
	return &repo{DB}
}

func (r *repo) FindAllProducts() ([]Product, error) {
	products := []Product{}
	if err := r.conn.Preload("Components").
		Preload("Components.License").
		Preload("UsageTypes").
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repo) FindProductByID(id int) (Product, error) {
	product := Product{}
	if err := r.conn.Preload("Components").
		Preload("Components.License").
		Preload("UsageTypes").
		Find(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *repo) SaveProduct(prod *Product) (Product, error) {
	if err := r.conn.Create(&prod).Error; err != nil {
		return *prod, err
	}

	return *prod, nil
}
