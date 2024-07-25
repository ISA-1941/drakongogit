package main

//import "math"
import "fmt"

/*
   Go program for
   Minimum number of edges between two vertices of a Graph
*/
type AjlistNode struct {
	// Vertices node key
	id   int
	next *AjlistNode
}

func getAjlistNode(id int) *AjlistNode {
	// return new AjlistNode
	return &AjlistNode{
		id,
		nil,
	}
}

type Vertices struct {
	data int
	next *AjlistNode
	last *AjlistNode
}

func getVertices(data int) *Vertices {
	// return new Vertices
	return &Vertices{
		data,
		nil,
		nil,
	}
}

type Graph struct {
	// Number of Vertices
	size   int
	result int
	node   []*Vertices
}

var k int = 0
var vs string

func getGraph(size int) *Graph {
	// return new Graph
	var me *Graph = &Graph{size, 0, make([]*Vertices, size)}
	me.setData()
	return me
}

// Set initial node value
func (graph *Graph) setData() {
	if graph.size <= 0 {
		fmt.Println("\nEmpty Graph")
	} else {
		for index := 0; index < graph.size; index++ {
			// Set initial node value
			graph.node[index] = getVertices(index)
		}
	}
}
func (graph *Graph) connection(start, last int) {
	// Safe connection
	var edge *AjlistNode = getAjlistNode(last)
	if graph.node[start].next == nil {
		graph.node[start].next = edge
	} else {
		// Add edge at the end
		graph.node[start].last.next = edge
	}
	// Get last edge
	graph.node[start].last = edge
}

//  Handling the request of adding new edge
func (graph *Graph) addEdge(start, last int) {
	if start >= 0 && start < graph.size && last >= 0 && last < graph.size {
		graph.connection(start, last)
	} else {
		// When invalid nodes
		fmt.Println("\nHere Something Wrong")
	}
}
func (graph Graph) printGraph() {
	if graph.size > 0 {
		// Print graph ajlist Node value
		for index := 0; index < graph.size; index++ {
			fmt.Print("\nAdjacency list of vertex ", index, " :")
			var edge *AjlistNode = graph.node[index].next
			for edge != nil {
				// Display graph node
				fmt.Print("  ", graph.node[edge.id].data)
				// Visit to next edge
				edge = edge.next
			}
		}
	}
}
func (graph *Graph) findMinimumEdge(start int, last int, visit []bool, length int) {
	k++
	fmt.Println("FME k = ", k, "length (+1) = ", length)
	f1 := start >= graph.size || last >= graph.size
	f2 := start < 0 || last < 0 || graph.size <= 0
	if f1 || f2 {
		return
	}
	if visit[start] == true {
		vs = "visit[start] == true"
		fmt.Println("\n if visit[start] == true --> Return --> k =", k, "  start = ", start, " ::", vs)
		return
	}
	// Here length are indicate number of edge
	if start == last && length < graph.result {
		// When new result
		fmt.Println("\n  if(start == last && length < graph.result) : k = ", k, "start == last ", start, last, length, graph.result)
		graph.result = length
	}
	// Here modified the value of visited node
	visit[start] = true
	// This is used to iterate nodes edges
	var edge *AjlistNode = graph.node[start].next
	fmt.Print("\n     k=", k, "   :", start, "-", last, ": length=", length, ": graph.result=", graph.result, visit[start])
	fmt.Print("\n        edge = graph.node[start].next --> ", edge, " : ", edge.next, " > ", visit[start])
	for edge != nil {
		fmt.Println("\n  ====================== edge != nil --> ===================")
		fmt.Println("Before k =", k, " edge.id = ", edge.id, " :  length = ", length)
		//fmt.Print("\n проход по узлам : edge.id, last, visit, length + 1 --> ", edge.id, last, visit, length + 1)
		graph.findMinimumEdge(edge.id, last, visit, length+1)
		fmt.Print("\n")
		// Visit to next edge
		edge = edge.next
		fmt.Print(" After --> k = ", k, "   : start =", start, " : last = ", last, " : visit = ", visit, " : length = ", length)
	}
	// Reset the value of visited node status
	visit[start] = false
}
func (graph Graph) minEdges(u, v int) {
	if graph.size <= 0 {
		// Empty graph
		return
	}
	// Auxiliary space which is used to store
	// information about visited node.
	// Set initial visited node status false
	var visit = make([]bool, graph.size)
	fmt.Println("\n")
	fmt.Println("\nminiEdges: visit = ", visit)
	graph.result = 1000
	//math.MaxInt64
	// Count edge between u to v
	graph.findMinimumEdge(u, v, visit, 0)
	// Display result
	fmt.Println("\nMinimum edges in between [",
		u, ",", v, "] is ", graph.result)
}
func main() {
	// 7 implies the number of nodes in graph
	var g *Graph = getGraph(7)
	// Connect node with an edge
	// First and second parameter indicate node
	g.addEdge(0, 1)
	g.addEdge(0, 6)
	g.addEdge(1, 2)
	g.addEdge(1, 4)
	g.addEdge(2, 0)
	g.addEdge(2, 4)
	g.addEdge(3, 5)
	g.addEdge(4, 3)
	g.addEdge(5, 4)
	g.addEdge(6, 3)
	fmt.Println("g--> ", g)
	// Print graph element
	g.printGraph()
	// Test
	// 2 is source node
	// 5 is destination node
	g.minEdges(2, 5)
}
