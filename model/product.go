package model

// Product represents a software product and its main properties.
// It is the root element in the product architecture model.
type Product struct {
	Name          string
	Description   string
	Version       string
	VCS           string
	ClearingState interface{} // TODO: Specify type
	RootDep       []*Dependency
	Components    []Component
	Infrastructure
	UsageType
}

// License represents a open source license.
type License struct {
	SPDXID string
}

// Copyright represents a copyright statement.
type Copyright struct{}

// Policy represents user/company specific rules for the usage of components and its licenses.
type Policy struct{}

// Artifact represents a digital artifact like source code files or binaries.
type Artifact struct {
	Path   string
	Hashes []interface{} // TODO: Specify type
}

// Deliverable represents a subset a product's components that are deliverable.
type Deliverable struct{} // deployment unit

// Infrastructure represents a dependency to required infrastructure, e.g. a runtime enviroment like the Java Runtime in a specific version.
type Infrastructure struct{}

// Vulnerability represents known security vulnerabilities of software components, e.g. CVE
type Vulnerability struct{}

// UsageType represents the scenario in which a product ist used, e.g. cloud service or internal usage only.
type UsageType interface{}
