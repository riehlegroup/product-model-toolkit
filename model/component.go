package model

import "fmt"

// Component represents a unit of compositon of the product, e.g. class, lib, module.
type Component struct {
	Name    string
	Pkg     string
	Version string
	License
	Copyright
	Artifact
}

// CmpID represents an component identifier which is a combination of package + name + version.
type CmpID string

// ID returns an identifier which is a combination of package + name + version.
func (c *Component) ID() CmpID {
	id := fmt.Sprintf("%v:%v:%v", c.Pkg, c.Name, c.Version)
	return CmpID(id)
}
