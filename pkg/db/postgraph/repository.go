// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package postgraph

import (
	"context"
	"errors"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/osrgroup/product-model-toolkit/model"
)

var (
	// ErrNotFound if a entity couldn't be found in the db.
	ErrNotFound = errors.New("Entity not found")
)

type repo struct {
	client *graphql.Client
}

// NewRepo create an instance for Postgraphile DB repository
func NewRepo(URL string) *repo {
	return &repo{graphql.NewClient(URL)}
}

// FindAllProducts returns all Products from the DB.
func (r *repo) FindAllProducts() ([]model.Product, error) {
	req := graphql.NewRequest(`
	query AllProds {
		allProducts {
		  nodes {
			id
			createdAt
			name
			vcs
			version
			description
			homepageUrl
			externalRef
			comment
		  }
		}
	  }
	`)

	var res struct {
		AllProducts struct {
			Nodes []model.Product
		}
	}

	if err := r.client.Run(context.Background(), req, &res); err != nil {
		return nil, err
	}

	return res.AllProducts.Nodes, nil
}

// FindProductByID return the Product with the given ID from the DB.
func (r *repo) FindProductByID(id int) (model.Product, error) {
	req := graphql.NewRequest(`
	query Prod($id: Int!) {
		productById(id: $id) {
			id
			createdAt
			name
			vcs
			version
			description
			homepageUrl
			externalRef
			comment
		}
	  }`)

	req.Var("id", id)

	var res struct {
		ProductByID model.Product
	}

	if err := r.client.Run(context.Background(), req, &res); err != nil {
		return model.Product{}, ErrNotFound
	}

	return res.ProductByID, nil
}

// SaveProduct store a Prodact to the DB.
func (r *repo) SaveProduct(prod *model.Product) (model.Product, error) {
	req := graphql.NewRequest(`
	mutation ProdMutation ($name: String!, $version: String!, $vcs: String, $comment: String, $description: String, $externalRef: String, $homepageUrl: String) {
		__typename
		createProduct(input: {product: {
			name: $name,
			version: $version,
			vcs: $vcs,
			description: $description, 
			externalRef: $externalRef, 
			homepageUrl: $homepageUrl, 
			comment: $comment}}) {
		  product {
			id
			createdAt
			name
			vcs
			version
			description
			homepageUrl
			externalRef
			comment
		  }
		}
	  }
	  `)

	req.Var("name", prod.Name)
	req.Var("version", prod.Version)
	req.Var("vcs", prod.VCS)
	req.Var("comment", prod.Comment)
	req.Var("description", prod.Description)
	req.Var("externalRef", prod.ExternalReference)
	req.Var("homepageUrl", prod.HomepageURL)

	var res struct {
		CreateProduct struct {
			Product model.Product
		}
	}

	if err := r.client.Run(context.Background(), req, &res); err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	fmt.Printf("Created product with id %v", res.CreateProduct.Product.ID)
	return res.CreateProduct.Product, nil
}

// DeleteProduct deletes a product with the given ID from the DB
func (r *repo) DeleteProduct(id int) error {
	req := graphql.NewRequest(`
	mutation DelMutation($id: Int!) {
		__typename
		deleteProductById(input: {id: $id}) {
		  product {
			id
		  }
		}
	  }
	  `)

	req.Var("id", id)

	var res interface{}

	if err := r.client.Run(context.Background(), req, &res); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Deleted product with id %v", id)
	return nil
}
