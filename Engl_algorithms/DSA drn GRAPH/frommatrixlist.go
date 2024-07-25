package main
import "fmt"
/*
    Go program for
    Construct adjacency list using adjacency matrix
*/
type Graph struct {
	vertices int
	adgeList [][]int
}
func getGraph(matrix [][]int ) * Graph {
    var me *Graph = &Graph {}
    me.vertices = len(matrix)
    me.adgeList = make([][]int,me.vertices)
    for i := 0 ; i < me.vertices ; i++ {
        me.adgeList = append(me.adgeList, make([]int,0))
    }
    me.makeAdjacencyList(matrix)
    return me
}
// Convert into adjacency list
func(this *Graph) makeAdjacencyList(matrix[][] int) {
	for i := 0 ; i < this.vertices ; i++ {
		for j := 0 ; j < this.vertices ; j++ {
			if matrix[i][j] == 1 {
				this.addEdge(i, j)
			}
		}
	}
}
func(this *Graph) addEdge(u, v int) {
	if u < 0 || u >= this.vertices || 
      v < 0 || v >= this.vertices {
		return
	}
	// Add node edge
	this.adgeList[u] = append(this.adgeList[u], v)
}
// Display graph nodes and edges
func(this Graph) printGraph() {
	fmt.Print("\n Graph Adjacency List ")
	for i := 0 ; i < this.vertices ; i++ {
		fmt.Print(" \n [", i, "] :")
		// iterate edges of i node
		for j := 0 ; j < len(this.adgeList[i]) ; j++ {
			if j != 0 {
				fmt.Print(" → ")
			}
			fmt.Print(" ", this.adgeList[i][j])
		}
	}
}
func main() {
	var matrix = [][] int {
		{0,1,1,0,1},
		{1,0,1,0,1},
		{1,1,0,1,0},
		{0,1,0,0,1},
		{1,1,0,1,0},
	}
	var g * Graph = getGraph(matrix)
	// Display graph element
	g.printGraph()
}