package main
import "fmt"
/*
    Go Program for 
    Directed graph representation using adjacency list
*/
type ListVertAdj struct {
    // Vertices node key
    id int
    next * ListVertAdj
}
func getListVertAdj(id int) * ListVertAdj {
    return &ListVertAdj {id,nil}
}
type Vertices struct {
    data int
    next * ListVertAdj
    last * ListVertAdj
}
func getVertices(data int) * Vertices {
    return &Vertices {data, nil, nil}
}
type Graph struct {
    // Number of Vertices
    size int
    node []*Vertices
}
func getGraph(size int) * Graph {
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
// Connect two nodes
func(this *Graph) connect(start, last int) {
	var edge * ListVertAdj = getListVertAdj(last)
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
    if start >= 0 && start < this.size && 
        last >= 0 && last < this.size {
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
            var temp * ListVertAdj = this.node[index].next
            for (temp != nil) {
                // Display graph node 
                fmt.Print(" temp = ", temp)
                // visit to next edge
                temp = temp.next
            }
        }
    }
}
func main() {
    // 5 implies the number of nodes in graph
    var g * Graph = getGraph(5)
    // Connect node with an edge      0,1  2,0  2,1
    g.addEdge(0, 1)
    g.addEdge(2, 0)
    g.addEdge(2, 1)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
    // print graph element
    g.printGraph()
}