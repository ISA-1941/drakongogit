package main

import "fmt"

/*
   Go program for
   DFS traversal of directed graph using recursion
*/
type AjlistNode struct {
	// Vertices node key
	id   int
	next *AjlistNode
}

func getAjlistNode(id int) *AjlistNode {
	return &AjlistNode{id, nil}
}

type Vertices struct {
	data int
	next *AjlistNode
	last *AjlistNode
}

var dfs, printDFS, printGraph int = 0, 0, 0

func getVertices(data int) *Vertices {
	return &Vertices{data, nil, nil}
}

type Graph struct {
	// Number of Vertices
	size int
	node []*Vertices
}

func getGraph(size int) *Graph {
	var me *Graph = &Graph{size, make([]*Vertices, size)}
	me.setData()
	return me
}

// Set initial node value
func (this *Graph) setData() {
	if this.size <= 0 {
		fmt.Println("\nEmpty Graph")
	} else {
		for index := 0; index < this.size; index++ {
			// Set initial node value
			this.node[index] = getVertices(index)
		}
	}
}

// Connect two nodes
func (this *Graph) connect(start, last int) {
	var edge *AjlistNode = getAjlistNode(last)
	if this.node[start].next == nil {
		this.node[start].next = edge
	} else {
		// Add edge at the end
		this.node[start].last.next = edge
	}
	// Get last edge
	this.node[start].last = edge
}

//  Handling the request of adding new edge
func (this *Graph) addEdge(start, last int) {
	if start >= 0 && start < this.size && last >= 0 &&
		last < this.size {
		// Safe connection
		this.connect(start, last)
	} else {
		// When invalid nodes
		fmt.Println("\nHere Something Wrong")
	}
}
func(this Graph) printGraph() {
    if this.size > 0 {
        // Print graph ajlist Node value
        for index := 0 ; index < this.size ; index++ {
            fmt.Print("\nAdjacency list of vertex ", index, " :")
            var temp * AjlistNode = this.node[index].next
            for (temp != nil) {
                // Display graph node 
                fmt.Print("  ", this.node[temp.id].data)
                // visit to next edge
                temp = temp.next
            }
        }
    }
}

// Dfs traversal of given node point
func (this *Graph) dfs(visit []bool, point int) {
	dfs++
	//fmt.Println(" \n dfs=", dfs, " ->", point, visit)
	if visit[point] {
		// We already visited this node
		return
	}
	// вершина посещена и стала активной
	visit[point] = true
	// Display visiting node point
	fmt.Print("              point of dfs = ", point)
	// Get node edge
	var temp *AjlistNode = this.node[point].next
	fmt.Println("dfs -->", this.node[point].data)
	// Iterate all edges in given node point
	for temp != nil {
		this.dfs(visit, temp.id)
		// next edge
		temp = temp.next
	}
}

// Handles the request of printing dfs sequence
func (this Graph) printDFS(point int) {
	printDFS++
	//fmt.Println("printDFS = ", printDFS)
	if this.size <= 0 || point < 0 || point >= this.size {
		return
	}
	// This is node visitor
	var visit = make([]bool, this.size)
	// There is no initial node visiting
	for i := 0; i < this.size; i++ {
		visit[i] = false
	}
	// Start traversal
	this.dfs(visit, point)
}
func main() {
	// 6 implies the number of nodes in graph
	var g *Graph = getGraph(6)
	/*
	g.addEdge(0, 1)
    g.addEdge(0, 4)
    g.addEdge(1, 2)
    g.addEdge(1, 3)
    g.addEdge(2, 3)
    g.addEdge(3, 4)
    g.addEdge(4, 5)
	*/
	g.addEdge(0, 1)
    g.addEdge(0, 5)
    g.addEdge(1, 1)
    g.addEdge(2, 1)
    g.addEdge(3, 0)
    g.addEdge(3, 3)
    g.addEdge(4, 2)
    g.addEdge(4, 3)
    g.addEdge(5, 1)
	g.printGraph()
	fmt.Print("\nDFS :\n")
	g.printDFS(4)
}

