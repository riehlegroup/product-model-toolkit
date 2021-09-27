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
	if _, ok := g.Vertices[data]; !ok {
		g.Vertices[data] = NewVertex(data)
	}
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