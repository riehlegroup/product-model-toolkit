// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

// Product represents a software product and its main properties.
// It is the root element in the product architecture model.
type Product struct {
	ID                int         `json:"id"`
	Name              string      `json:"name,omitempty"`
	Version           string      `json:"version,omitempty"`
	VCS               string      `json:"vcs,omitempty"`
	Description       string      `json:"description,omitempty"`
	Comment           string      `json:"comment,omitempty"`
	HomepageURL       string      `json:"homepageUrl,omitempty"`
	ExternalReference string      `json:"externalRef,omitempty"`
	ClearingState     interface{} `json:"-"` // TODO: Specify type
	Components        []Component `json:"components"`
	DepGraph          DepGraph    `json:"-"`
	Infrastructure    `json:"-"`
	UsageTypes        []UsageType `json:"usageTypes,omitempty"`
}

// Copyright represents a copyright statement.
type Copyright struct{}

// Policy represents user/company specific rules for the usage of components and its licenses.
type Policy struct{}

// Deliverable represents a subset a product's components that are deliverable.
type Deliverable struct{} // deployment unit

// Infrastructure represents a dependency to required infrastructure, e.g. a runtime environment like the Java Runtime in a specific version.
type Infrastructure struct{}

// Vulnerability represents known security vulnerabilities of software components, e.g. CVE
type Vulnerability struct{}
