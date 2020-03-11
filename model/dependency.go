package model

// Dependency represents the relationship between two components.
type Dependency struct {
	Type DepType
	From *Component
	To   *Component
}

// Next returns true if dependency points to a component.
func (d *Dependency) Next() bool {
	return d.To != nil
}

// DepType represents the type of a dependency, e.g. static linking.
type DepType string

const (
	StaticLinking  DepType = "STATIC_LINKING"
	DynamicLinking DepType = "DYNAMIC_LINKING"
)

type Iterator interface {
	Next() bool
}
