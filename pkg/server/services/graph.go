package services

import "fmt"

type Vertex struct {
	Data string

	Vertices map[string]*Vertex
}

func NewVertex(data string) *Vertex {
	return &Vertex{
		Data: data,
		Vertices: make(map[string]*Vertex),
	}
}

type Graph struct {
	Vertices map[string]*Vertex
	directed bool
}

func NewDirectedGraph() *Graph {
return &Graph{
	Vertices: make(map[string]*Vertex),
	directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}


func (g *Graph) AddVertex(data string) {
	v := NewVertex(data)
	g.Vertices[data] = v
}

func (g *Graph) AddEdge(from, to string) error {

	v1:= g.Vertices[from]
	v2:= g.Vertices[to]

	if v1 == nil || v2 == nil {
		return fmt.Errorf("Vertex not found")
	}

	if _, ok := v1.Vertices[v2.Data]; ok {
		return fmt.Errorf("Edge already exists")
	}

	v1.Vertices[v2.Data] = v2
	if !g.directed && v1.Data != v2.Data {
		v2.Vertices[v1.Data] = v1
	}

	g.Vertices[v1.Data] = v1
	g.Vertices[v2.Data] = v2

	return nil
}


func (g *Graph) RemoveEdge(from, to string) error {
	v1:= g.Vertices[from]
	v2:= g.Vertices[to]

	if v1 == nil || v2 == nil {
		return fmt.Errorf("Vertex not found")
	}

	if _, ok := v1.Vertices[v2.Data]; !ok {
		return fmt.Errorf("Edge not found")
	}

	delete(v1.Vertices, v2.Data)
	if !g.directed && v1.Data != v2.Data {
		delete(v2.Vertices, v1.Data)
	}

	return nil
}

func IsAncestor(g *Graph, ancestor, descendant string) bool {
	if g.directed {
		return isAncestorDirected(g, ancestor, descendant)
	} else {
		return isAncestorUndirected(g, ancestor, descendant)
	}
}

func isAncestorDirected(g *Graph, ancestor, descendant string) bool {
	if g.Vertices[ancestor] == nil || g.Vertices[descendant] == nil {
		return false
	}

	if g.Vertices[ancestor].Data == descendant {
		return true
	}

	for _, v := range g.Vertices[ancestor].Vertices {
		if v.Data == descendant {
			return true
		}
	}

	return false
}

func isAncestorUndirected(g *Graph, ancestor, descendant string) bool {
	if g.Vertices[ancestor] == nil || g.Vertices[descendant] == nil {
		return false
	}

	if g.Vertices[ancestor].Data == descendant {
		return true
	}

	for _, v := range g.Vertices[ancestor].Vertices {
		if v.Data == descendant {
			return true
		}
	}

	for _, v := range g.Vertices[descendant].Vertices {
		if v.Data == ancestor {
			return true
		}
	}

	return false
}
