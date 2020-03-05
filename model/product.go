package model

// Product represents a software product and its main properties. 
// It is the root element in the product architecture model.
type Product struct {}

// Component represents a unit of compositon of the product, e.g. class, lib, module.
type Component struct {}

// Dependency represents the relationship between two components.
type Dependency struct {}

// License represents a open source license.
type License struct {}

// Copyright represents a copyright statement.
type Copyright struct {}

// Policy represents user/company specific rules for the usage of components and its licenses.
type Policy struct {}

// Artifact represents a digital artifact like source code files or binaries.
type Artifact struct {}

// Deliverable represents a subset a product's components that are deliverable.
type Deliverable struct {} // deployment unit

// Infrastructure represents a dependency to required infrastructure, e.g. a runtime enviroment in a specific version.
type Infrastructure struct {} // Inf. dep. e.g. Java Runtime

// Vulnerability represents known security vulnerabilities of software components.
type Vulnerability struct {} // e.g. CVE
