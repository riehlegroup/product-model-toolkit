package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	pb "pmt/model"
)


type Product struct {
	ID             string     `json:"id"`
	Version        string     `json:"version"`
	Vcs            string     `json:"vcs"`
	Description    string     `json:"description"`
	Comment        string     `json:"commenet"`
	HomePageURL    string     `json:"home_page_url"`
	ExternalRef    string     `json:"external_ref"`
	Components     Components `json:"components"`
	UsageTypes     string     `json:"usage_types"`
	ClearingState  string     `json:"clearing_state"`
	DepGraph       string     `json:"dep_graph"`
	Infrastructure string     `json:"infrastructure"`
}


type Component struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Package string  `json:"package"`
	Version string  `json:"version"`
	License License `json:"license"`
}

type Components []*Component

type License struct {
	ID         string `json:"id"`
	DeclaredLicense string `json:"declared_license"`
	ConcludedLicense string `json:"concluded_license"`
}

func MarshalLicenseCollection(license *pb.License) *License {
	collection := MarshalLicense(license)
	return collection
}

func UnmarshalLicenseCollection(license *License) *pb.License {
	collection := UnmarshalLicense(license)
	return collection
}

func MarshalComponentCollection(components []*pb.Component) []*Component {
	collection := make([]*Component, 0)
	for _, component := range components {
		collection = append(collection, MarshalComponent(component))
	}
	return collection
}

func UnmarshalComponentCollection(components []*Component) []*pb.Component {
	collection := make([]*pb.Component, 0)
	for _, component := range components {
		collection = append(collection, UnmarshalComponent(component))
	}
	return collection
}

func UnmarshalProductCollection(products []*Product) []*pb.Product {
	collection := make([]*pb.Product, 0)
	for _, product := range products {
		collection = append(collection, UnmarshalProduct(product))
	}
	return collection
}

func UnmarshalLicense(license *License) *pb.License {
	return &pb.License{
		Id:               license.ID,
		DeclaredLicense:  license.DeclaredLicense,
		ConcludedLicense: license.ConcludedLicense,
	}
}

func MarshalLicense(license *pb.License) *License {
	return &License{
		ID:              license.Id,
		DeclaredLicense:  license.DeclaredLicense,
		ConcludedLicense: license.ConcludedLicense,
	}
}

func UnmarshalComponent(component *Component) *pb.Component {
	return &pb.Component{
		Id:      component.ID,
		Name:    component.Name,
		Package: component.Package,
		Version: component.Version,
		License: UnmarshalLicenseCollection(&component.License),
	}
}


func MarshalComponent(component *pb.Component) *Component {
	license := MarshalLicenseCollection(component.License)
	return &Component{
		ID:      component.Id,
		Name:    component.Name,
		Package: component.Package,
		Version: component.Version,
		License: *license,
	}
}



func MarshalProduct(product *pb.Product) *Product {
	components := MarshalComponentCollection(product.Components)
	return &Product{
		ID:             product.Id,
		Version:        product.Version,
		Vcs:            product.Vcs,
		Description:    product.Description,
		Comment:        product.Comment,
		HomePageURL:    product.HomePageUrl,
		ExternalRef:    product.ExternalRef,
		Components:     components,
		UsageTypes:     product.UsageTypes,
		ClearingState:  product.ClearingState,
		DepGraph:       product.DepGraph,
		Infrastructure: product.Infrastructure,
	}
}

func UnmarshalProduct(product *Product) *pb.Product {
	return &pb.Product{
		Id:             product.ID,
		Version:        product.Version,
		Vcs:            product.Vcs,
		Description:    product.Description,
		Comment:        product.Comment,
		HomePageUrl:    product.HomePageURL,
		ExternalRef:    product.ExternalRef,
		Components:     UnmarshalComponentCollection(product.Components),
		UsageTypes:     product.UsageTypes,
		ClearingState:  product.ClearingState,
		DepGraph:       product.DepGraph,
		Infrastructure: product.Infrastructure,
	}
}

type repository interface {
	Create(ctx context.Context, product *Product) error
	GetAll(ctx context.Context) ([]*Product, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) Create(ctx context.Context, product *Product) error {
	_, err := repository.collection.InsertOne(ctx, product)
	return err
}

func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Product, error) {
	cur, err := repository.collection.Find(ctx, nil, nil)
	var products []*Product
	for cur.Next(ctx) {
		var product *Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}
