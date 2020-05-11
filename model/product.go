// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

// Product represents a software product and its main properties.
// It is the root element in the product architecture model.
type Product struct {
	ID             int           `json:"id"`
	Name           string        `json:"name,omitempty"`
	Info           Info          `json:"info,omitempty"`
	Version        string        `json:"version,omitempty"`
	VCS            string        `json:"vcs,omitempty"`
	ClearingState  interface{}   `json:"-"` // TODO: Specify type
	RootDep        []*Dependency `json:"-"`
	Components     []Component   `json:"components"`
	Infrastructure `json:"-"`
	UsageTypes     []UsageType `json:"usage-types,omitempty"`
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

// Infrastructure represents a dependency to required infrastructure, e.g. a runtime environment like the Java Runtime in a specific version.
type Infrastructure struct{}

// Vulnerability represents known security vulnerabilities of software components, e.g. CVE
type Vulnerability struct{}
