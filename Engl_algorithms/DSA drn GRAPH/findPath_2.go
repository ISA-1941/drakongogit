package main
import "strconv"
import "fmt"
/*
    Go program for 
    Print all path between given vertices in a directed graph
*/
type AjlistNode struct {
	// Vertices node key
	id int
	next * AjlistNode
}
func getAjlistNode(id int) * AjlistNode {
	// return new AjlistNode
	return &AjlistNode {
		id,
		nil,
	}
}
type Vertices struct {
	data int
	next * AjlistNode
	last * AjlistNode
}
func getVertices(data int) * Vertices {
	// return new Vertices
	return &Vertices {
		data,
		nil,
		nil,
	}
}
type Graph struct {
	// Number of Vertices
	size int
	node []*Vertices
}
func getGraph(size int) * Graph {
	// return new Graph
	var me *Graph = &Graph {size,make([]*Vertices, size)}
    me.setData()
    return me
}
// Set initial node value
func(this *Graph) setData() {
	if this.size <= 0 {
		fmt.Println("\nEmpty Graph")
	} else {
		for index := 0 ; index < this.size ; index++ {
			// Set initial node value
			this.node[index] = getVertices(index)
		}
	}
}
func(this *Graph) connection(start, last int) {
	// Safe connection
	var edge * AjlistNode = getAjlistNode(last)
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
func(this *Graph) addEdge(start, last int) {
	if start >= 0 && start < this.size && last >= 0 && 
		last < this.size && this.size > 0 {
		this.connection(start, last)
	} else {
		// When invalid nodes
		fmt.Println("\nHere Something Wrong")
	}
}
func(this Graph) printGraph() {
	// Print graph ajlist Node value
	for index := 0 ; index < this.size ; index++ {
		fmt.Print("\nAdjacency list of vertex ", index, " :")
		var edge * AjlistNode = this.node[index].next
		for (edge != nil) {
			// Display graph node 
			fmt.Print("  ", this.node[edge.id].data)
			// Visit to next edge
			edge = edge.next
		}
	}
}
func(this Graph) findPath(start int, last int, 
	path string, visit[] bool) {
	if  start >= this.size || last >= this.size || 
		start < 0 || last < 0 || this.size <= 0 {
		return
	}
	if visit[start] == true {
		// When have already visited
		return
	}
	if start == last {
		// Display result
		fmt.Println("(", path, ")")
	}
	// Here modified the value of visited node
	visit[start] = true
	// This is used to iterate nodes edges
	var edge * AjlistNode = this.node[start].next
	for (edge != nil) {
		this.findPath(edge.id, last, 
			path + " " + strconv.Itoa(edge.id), visit)
		// Visit to next edge
		edge = edge.next
	}
	// Reset the value of visited node status
	visit[start] = false
}
// Print all paths between two vertices in directed graph
func(this Graph) allPath(start, last int) {
	if this.size <= 0 {
		return
	}
	// Use to track node
	var visit = make([] bool, this.size)
	// Set initial value
	for i := 0 ; i < this.size ; i++ {
		visit[i] = false
	}
	fmt.Println("\nResult : ")
	this.findPath(start, last, ""+ strconv.Itoa(start), visit)
}
func  main()  {
    // item 4
    var graph *Graph = getGraph(7)
    // item 5
    	graph.addEdge(0, 1)
    	graph.addEdge(0, 6)
    	graph.addEdge(1, 2)
    	graph.addEdge(1, 4)
    	graph.addEdge(2, 0)
    	graph.addEdge(2, 4)
    	graph.addEdge(3, 5)
    	graph.addEdge(4, 3)
    	graph.addEdge(5, 4)
    	graph.addEdge(6, 3)
    // item 6
    graph.printGraph()
    // item 7
    graph.findPath(1, 4)
}