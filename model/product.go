// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Product represents a software product and its main properties.
// It is the root element in the product architecture model.
type Product struct {
	gorm.Model       `json:"-"`
	Name              string      `gorm:"column:name;not null" json:"name,omitempty"`
	Version           string      `gorm:"column:version" json:"version,omitempty"`
	License           string      `gorm:"column:license" json:"license,omitempty"` // license added for checking the compatibiltiy
	VCS               string      `gorm:"column:vcs" json:"vcs,omitempty"`
	Description       string      `gorm:"column:description" json:"description,omitempty"`
	Comment           string      `gorm:"column:comment" json:"comment,omitempty"`
	HomepageURL       string      `gorm:"column:home_page_url" json:"home_page_url,omitempty"`
	ExternalReference string      `gorm:"column:external_reference" json:"external_reference,omitempty"`
	ClearingState     string      `json:"-"`
	Components        []Component `gorm:"foreignKey:ProductRefer" json:"components"`
	DepGraphRefer     int         `json:"-"`
	DepGraph          DepGraph    `gorm:"foreignKey:DepGraphRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Infrastructure    string      `json:"-"`
	UsageTypes        []UsageType `gorm:"foreignKey:ProductRefer" json:"usageTypes,omitempty"`
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

// Component represents a unit of composition of the product, e.g. class, lib, module.
type Component struct {
	gorm.Model    `json:"-"`
	ProductRefer  uint 	`json:"-"`
	UID           string   `gorm:"column:uid;not null" json:"uid"`
	Name          string   `gorm:"column:name" json:"name"`
	Package       string   `gorm:"column:package" json:"package"`
	Version       string   `gorm:"column:version" json:"version"`
	LicenseRefer  int      `json:"-"`
	License       License  `gorm:"foreignKey:LicenseRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"license,omitempty"`
	Copyright     string   `json:"-"`
	ArtifactRefer int      `json:"-"`
	Artifact      Artifact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

// CmpID represents an component identifier which is a combination of package + name + version.
type CmpID string

// ID returns an identifier which is a combination of package + name + version.
func (c *Component) ID() CmpID {
	id := fmt.Sprintf("%v:%v:%v", c.Package, c.Name, c.Version)
	return CmpID(id)
}

// ContainsComp reports whether a component with the given component id is present in cmps
func ContainsComp(cmps []Component, cid CmpID) bool {
	for _, v := range cmps {
		if v.ID() == cid {
			return true
		}
	}

	return false
}

// License represents a open source license.
type License struct {
	gorm.Model		 `json:"-"`
	SPDXID           string `gorm:"column:spdx_id;not null" json:"spdxId,omitempty"`
	DeclaredLicense  string `gorm:"column:declared_licesne" json:"declaredLicense,omitempty"`
	ConcludedLicense string `gorm:"column:concluded_license" json:"concludedLicense,omitempty"`
}

func (l *License) toString() string {
	if l.SPDXID != "" {
		return l.SPDXID
	}

	if l.ConcludedLicense != "" {
		return l.ConcludedLicense
	}

	return l.DeclaredLicense
}

// Artifact represents a digital artifact like source code files or binaries.
type Artifact struct {
	gorm.Model
	Path      string `gorm:"column:path" json:"path"`
	Name      string `gorm:"column:name" json:"name"`
	IsDir     bool   `gorm:"column:is_dir" json:"isDir"`
	HashRefer int    `json:"-"`
	Hash      Hash   `gorm:"foreignKey:HashRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"hash"`
}

// Hash represents the hash value of a artifact.
type Hash struct {
	gorm.Model
	MD5    string `gorm:"column:md5" json:"md5,omitempty"`
	SHA1   string `gorm:"column:sha1" json:"sha1,omitempty"`
	SHA256 string `gorm:"column:sha256" json:"sha256,omitempty"`
}

// DepGraph represents a component dependency graph.
type DepGraph struct {
	// Dependencies map[CmpID]map[CmpID]Dependency `json:"dependencies, omitempty"`
	Dependencies string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// Dependency represents the relationship between two components.
type Dependency struct {
	gorm.Model
	From    string `gorm:"column:from" json:"from,omitempty"`
	To      string `gorm:"column:to" json:"to,omitempty"`
	Linking string `gorm:"column:linking" json:"linking,omitempty"`
}

// LinkingType represents the type of linking between dependencies, e.g. static linking.
type LinkingType string

const (
	// StaticLinked represents static linking between components
	StaticLinked LinkingType = "STATIC_LINKED"
	// DynamicLinked represents dynamic linking between components
	DynamicLinked LinkingType = "DYNAMIC_LINKED"
)

// String returns a string representation of a dependency.
func (d *Dependency) String() string {
	return fmt.Sprintf("(%s) -> (%s) [linking: '%s']", d.From, d.To, d.Linking)
}

// // AddDependency adds a new dependency between two components.
// func (g *DepGraph) AddDependency(from, to CmpID, linking LinkingType) {
// 	if _, ok := g.Deps[from]; !ok {
// 		g.Deps[from] = make(map[CmpID]Dependency)
// 	}

// 	g.Deps[from][to] = Dependency{From: from, To: to, Linking: linking}
// }

// Info represents generic additional information related to a product or component.
type Info struct {
	Description       string `json:"description,omitempty"`
	Comment           string `json:"comment,omitempty"`
	HomepageURL       string `json:"homepage-url,omitempty"`
	ExternalReference string `json:"external-ref,omitempty"`
}

// UsageType represents the scenario in which a product ist used, e.g. cloud service or internal usage only.
type UsageType struct {
	gorm.Model
	ProductRefer uint
	Name         string `gorm:"column:name" json:"name"`
}
