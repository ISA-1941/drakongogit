// Autogenerated with DRAKON Editor 1.31

package main

import (
	"fmt"
	"math"
	"strconv"
)

type Graph struct {
	size  int
	nodes map[int][]int
}

var mapLen = make(map[string]float64)
var mapPrice = make(map[string]float64)
var pass bool

func (graph *Graph) addEdge(start, last int) {
	// item 54
	if start >= 0 && start < graph.size &&
		last >= 0 && last < graph.size {
		// item 57
		graph.nodes[start] =
			append(graph.nodes[start], last)
	} else {
		// item 58
		fmt.Println("\nError")
	}
}

func (graph Graph) allPaths(start, last int) {
	// item 44
	visited := make(map[int]bool)
	for node := range graph.nodes {
		// item 47
		visited[node] = false
	}
	// item 48
	graph.findPaths(start, last,
		visited, []int{})
}

func (graph Graph) findPaths(start int, last int, visited map[int]bool, path []int) {
	// item 20
	fmt.Println(" --- FindPath --- FindPath ---")
	var lenWay, priceLen float64 = 0.0, 0.0
	// item 21
	visited[start] = true
	// item 22
	path = append(path, start)
	// item 23
	fmt.Println("start = ", start, "path = ", path, "visited[start] = ", visited[start])
	// item 24
	if start == last {
		// item 27
		pass = true
		// item 280001
		i := 0
		for {
			// item 280002
			if i < len(path)-1 {

			} else {
				break
			}
			// item 30
			key := strconv.Itoa(path[i]) +
				"-" + strconv.Itoa(path[i+1])
			lenWay = lenWay + mapLen[key]
			priceLen = priceLen +
				mapLen[key]*mapPrice[key]
			// item 280003
			i++
		}
		// item 31
		fmt.Println("\n     Length of travel | ",
			math.Round(lenWay*100)/100, "km")
		fmt.Println(" Cost of travel |",
			math.Round(priceLen*100)/100, "$")
		fmt.Println("path = ", path)
	} else {
		for _, neighbor := range graph.nodes[start] {
			// item 34
			if !visited[neighbor] {
				// item 37
				graph.findPaths(neighbor,
					last, visited, path)
			} else {

			}
		}
	}
	// item 38
	fmt.Println("\n           START = ", start, "path ДО ", path, "visited = ", visited)
	visited[start] = false
	path = path[:len(path)-1]
	fmt.Println("           path После ", path, "          visited = ", visited)
}

func NewGraph() *Graph {
	// item 64
	return &Graph{
		nodes: make(map[int][]int),
	}
}

func main() {
	// item 74
	graph := NewGraph()
	graph.size = 7
	pass = false
	// item 76
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
	// item 75
	mapLen["0-1"] = 10.45
	mapLen["0-6"] = 13.57
	mapLen["1-2"] = 18.83
	mapLen["1-4"] = 14.23
	mapLen["2-0"] = 43.67
	mapLen["2-4"] = 27.93
	mapLen["3-5"] = 47.33
	mapLen["4-3"] = 29.27
	mapLen["5-4"] = 57.69
	mapLen["6-3"] = 33.31
	// item 77
	mapPrice["0-1"] = 1.50
	mapPrice["0-6"] = 1.80
	mapPrice["1-2"] = 2.00
	mapPrice["1-4"] = 3.00
	mapPrice["2-0"] = 1.20
	mapPrice["2-4"] = 1.30
	mapPrice["3-5"] = 1.00
	mapPrice["4-3"] = 3.30
	mapPrice["5-4"] = 1.00
	mapPrice["6-3"] = 1.60
	// item 78
	p1 := 2 // Initial vertex of the path
	p2 := 5 // Final peak of the path
	graph.allPaths(p1, p2)
	// item 79
	if pass == false {
		// item 82
		fmt.Println("No path between ",
			p1, "and ", p2)
	} else {

	}
}
