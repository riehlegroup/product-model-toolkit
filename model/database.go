// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	git "gopkg.in/src-d/go-git.v4"
	"net/url"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() (*gorm.DB, error) {
	username := os.Getenv("POSTGRES_USER")
	dbPort := os.Getenv("POSTGRES_PORT")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string

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

// find all products
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

// find product by id
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

// delete product by id
func (r *repo) DeleteProductByID(id int) error {
	if err := r.conn.Delete(&Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) Download(downloadDetails []string) (*DownloadData, error) {
	u, path := downloadDetails[0], downloadDetails[1]

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      u,
		Progress: os.Stdout,
	})

	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return nil, err
	}

	downloadData := new(DownloadData)
	downloadData.Path = path
	tempUrl, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	downloadData.Url = u
	downloadData.Slug = string(tempUrl.Path)[1:]

	return downloadData, nil
}

// save product
func (r *repo) SaveProduct(prod *Product) (Product, error) {
	if err := r.conn.Create(&prod).Error; err != nil {
		return *prod, err
	}

	return *prod, nil
}

func (r *repo) StoreDownloadedRepo(data *DownloadData) error {
	if err := r.conn.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) FindAllDownloadedRepos() ([]DownloadData, error) {
	downloadedData := []DownloadData{}
	if err := r.conn.Find(&downloadedData).Error; err != nil {
		return nil, err
	}
	return downloadedData, nil
}
