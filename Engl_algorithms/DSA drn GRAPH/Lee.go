package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key int
	adjacent []*Vertex
}
// добавить вершину
func (g *Graph) AddVertex (k int) {
	if contains (g.vertices, k) {
		err :=  fmt.Errorf("Vertex %v not added becouse it is an existing key", k)
		fmt.Println(err.Error())
	} else {
	g.vertices = append(g.vertices, &Vertex {key: k})
	}
}

func contains (s []*Vertex, k int) bool{ 
	for _,v := range s {
		if k == v.key {
			return true
}
}
}

// Печать графа
func (g *Graph) print() {
	for _,v := range g.vertices {
		fmt.Print("\nvertex %v : ", v.key)
		for _,v := range v.adjacent {
			fmt.Print("%v", v.key)
		}
	}
}
// Создать ребро
func (g *Graph) addEdge(from, to int)  {
	fromVertex := g.getVertex(from)
	toVertex :=  g.getVertex(to)
}

func (g *Graph) getVertex(k int) *Vertex {
	for i,v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
}


func main () {
	test := &Graph {}
	for i := 0; i < 5;  i++ {
		test.AddVertex(i)
		}
		test.AddVertex(0)
		test.AddVertex(0)
		test.AddVertex(0)
		test.AddVertex(0)
		test.print()

}