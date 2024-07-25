package main
import "fmt"
/*
    Go Program for
    Find if there is a path of more than k length from a source
*/
type AjlistNode struct {
	// Vertices node key
	id int
	length int
	next * AjlistNode
}
func getAjlistNode(id int, length int) * AjlistNode {
	var me *AjlistNode = &AjlistNode {}
	// Set value of node key
	me.id = id
	me.length = length
	me.next = nil
	return me
}
type Vertices struct {
	data int
	next * AjlistNode
}
func getVertices(data int) * Vertices {
	var me *Vertices = &Vertices {}
	me.data = data
	me.next = nil
	return me
}
type Graph struct {
	// Number of Vertices
	size int
	node []*Vertices
}
func getGraph(size int) * Graph {
	var me *Graph = &Graph {}
	//set value
	me.size = size
	me.node = make([]*Vertices, size)
	me.setData()
	return me
}
// Set initial node value
func(this *Graph) setData() {
	if this.node == nil {
		fmt.Println("\nEmpty Graph")
	} else {
		for index := 0 ; index < this.size ; index++ {
			this.node[index] = getVertices(index)
		}
	}
}
// Connect two nodes
func(this *Graph) connect(start, last, length int) {
	var new_edge * AjlistNode = getAjlistNode(last, length)
	if this.node[start].next == nil {
		this.node[start].next = new_edge
	} else {
		var temp * AjlistNode = this.node[start].next
		for (temp.next != nil) {
			temp = temp.next
		}
		temp.next = new_edge
	}
}
// Add edge of two nodes
func(this *Graph) addEdge(start, last, length int) {
	if start >= 0 && start < this.size && last >= 0 && 
	last < this.size && this.node != nil {
		this.connect(start, last, length)
		if start == last {
			// Self looping situation
			return
		}
		this.connect(last, start, length)
	} else {
		fmt.Println("\nHere Something Wrong")
	}
}
func(this Graph) printGraph() {
	if this.size > 0 && this.node != nil {
		// Print graph ajlist Node value
		for index := 0 ; index < this.size ; index++ {
			fmt.Print("\nAdjacency list of vertex ", index, " :")
			var temp * AjlistNode = this.node[index].next
			for (temp != nil) {
				fmt.Print(" ", this.node[temp.id].data)
				// visit to next edge
				temp = temp.next
			}
		}
	}
}
func(this Graph) isLengthGreaterThanK(
	index int, visit[] bool, sum int, k int) bool {
	//	fmt.Println("index = ", index, "sum = ", sum)
	if index < 0 || index >= this.size {
		return false
	}
	if visit[index] == true {
		// When node is already includes in current path
		return false
	}
	if sum > k {
		// When length sum is greater than k
		return true
	}
	// Here modified  the value of visited node
	visit[index] = true
	// This is used to iterate nodes edges
	var temp * AjlistNode = this.node[index].next
	for (temp != nil) {
		if this.isLengthGreaterThanK(temp.id, 
			visit, sum + (temp.length), k) {
			// Found path with length greater than k
			return true
		}
		// Visit to next edge
		temp = temp.next
	}
	// Reset the value of visited node status
	visit[index] = false
	return false
}
func(this Graph) checkPathGreaterThanK(source, k int) {
	// Set initial visited node status 
	var visit = make([] bool, this.size)
	fmt.Print("\n Source node : ", source)
	fmt.Print("\n Length  k : ", k)
	if this.isLengthGreaterThanK(source, visit, 0, k) {
		fmt.Print("\n Result : YES\n")
	} else {
		fmt.Print("\n Result : NO\n")
	}
}
func main() {
	// 10 implies the number of nodes in graph
	var g * Graph = getGraph(10)
	g.addEdge(0, 1, 9)
	g.addEdge(0, 2, 10)
	g.addEdge(0, 9, 4)
	g.addEdge(1, 3, 6)
	g.addEdge(1, 9, 5)
	g.addEdge(2, 4, 3)
	g.addEdge(2, 7, 7)
	g.addEdge(2, 9, 2)
	g.addEdge(3, 5, 5)
	g.addEdge(3, 7, 4)
	g.addEdge(3, 9, 3)
	g.addEdge(4, 6, 4)
	g.addEdge(5, 6, 3)
	g.addEdge(5, 7, 8)
	g.addEdge(6, 7, 8)
	g.addEdge(6, 8, 4)
	// print graph element
	g.printGraph()
	// Test A
	var source int = 0
	var k int = 48
	g.checkPathGreaterThanK(source, k)
	// Test B
	source = 0
	k = 51
	g.checkPathGreaterThanK(source, k)
	// Test C
	source = 8
	k = 54
	g.checkPathGreaterThanK(source, k)
}