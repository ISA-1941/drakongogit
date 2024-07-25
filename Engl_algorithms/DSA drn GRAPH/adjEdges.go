package main

import "fmt"

type Graph struct {
	nodes map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int][]int),
	}
}

func (g *Graph) addEdge(src, dest int) {
    fmt.Println("addEdge: src = ", src, "dest = ", dest)
	g.nodes[src] = append(g.nodes[src], dest)
	fmt.Println("addEdge: src = ", src, "g.nodes[src] = ", g.nodes[src])
}

func (g *Graph) printAdjacencyList() {
	for node, neighbors := range g.nodes {
		fmt.Printf("%d: ", node)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()

	graph.addEdge(0, 1)
	graph.addEdge(0, 5)
	graph.addEdge(0, 2)
	graph.addEdge(1, 1)
	graph.addEdge(2, 1)
	graph.addEdge(3, 0)
	graph.addEdge(3, 3)
	graph.addEdge(4, 2)
	graph.addEdge(4, 3)
	graph.addEdge(5, 1)

	graph.printAdjacencyList()
}
