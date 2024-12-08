// Autogenerated with DRAKON Editor 1.31

package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
    Go Program for
 Print all path between given vertices in a directed graph
*/
type AjlistNode struct {
	// Vertices node key
	id   int
	next *AjlistNode
}

type Vertices struct {
	data int
	next *AjlistNode
	last *AjlistNode
}

type Graph struct {
	// Number of Vertices
	size int
	node []*Vertices
}

var mapLen = make(map[string]float64)
var mapPrice = make(map[string]float64)

func (graph Graph) allPath(start int, last int) {
	// item 143
	if graph.size <= 0 {
		// item 146
		return
	} else {

	}
	// item 147
	var visit = make([]bool, graph.size)
	// item 1480001
	i := 0
	for {
		// item 1480002
		if i < graph.size {

		} else {
			break
		}
		// item 150
		visit[i] = false
		// item 1480003
		i++
	}
	// item 151
	fmt.Println("\n Результат : ")
	// item 152
	graph.findPath(start, last, ""+
		strconv.Itoa(start), visit)
}

func (graph Graph) printGraph() {
	// item 76
	if graph.size > 0 {
		// item 790001
		index := 0
		for {
			// item 790002
			if index < graph.size {

			} else {
				break
			}
			// item 81
			//fmt.Print("\n Список связности узла ",
			//               index, " : ")
			// item 82
			var temp *AjlistNode = graph.node[index].next
			for {
				// item 83
				if temp != nil {

				} else {
					break
				}
				// item 86
				//fmt.Print("  ", graph.node[temp.id].data)
				// item 87
				temp = temp.next
			}
			// item 790003
			index++
		}
	} else {

	}
}

func (graph *Graph) addEdge(start, last int) {
	// item 13
	if start >= 0 && start < graph.size &&
		last >= 0 && last < graph.size {
		// item 16
		var edge *AjlistNode = getAjlistNode(last)
		// item 17
		if graph.node[start].next == nil {
			// item 20
			graph.node[start].next = edge
		} else {
			// item 21
			graph.node[start].last.next = edge
		}
		// item 23
		graph.node[start].last = edge
	} else {
		// item 22
		fmt.Println("\nHere Something Wrong")
	}
}

func (graph *Graph) connect(start, last int) {
	// item 63
	var edge *AjlistNode = getAjlistNode(last)
	// item 64
	if graph.node[start].next == nil {
		// item 66
		graph.node[start].next = edge
		fmt.Println("graph.node[start].next -->",
			graph.node[start].next)
	} else {
		// item 67
		graph.node[start].last.next = edge
		fmt.Println("graph.node[start].last.next -->",
			graph.node[start].last.next)
	}
	// item 68
	graph.node[start].last = edge
}

func (graph *Graph) setData() {
	// item 27
	if graph.size <= 0 {
		// item 30
		fmt.Println("\nEmpty Graph")
	} else {
		// item 310001
		index := 0
		for {
			// item 310002
			if index < graph.size {

			} else {
				break
			}
			// item 33
			graph.node[index] =
				getVertices(index)
			// item 310003
			index++
		}
	}
}

func (graph Graph) findPath(start int, last int, path string, visit []bool) {
	// item 293
	var lenWay, priceLen float64 = 0.0, 0.0
	// item 262
	f1 := start >= graph.size ||
		last >= graph.size
	// item 263
	f2 := start < 0 || last < 0 ||
		graph.size <= 0
	// item 260
	if f1 || f2 {
		// item 264
		return
	} else {

	}
	// item 297
	if visit[start] == true {
		// item 300
		return
	} else {

	}
	// item 267
	if start == last {
		// item 268
		fmt.Print("Вершины пути : (", path, ")")
		// item 2690001
		i := 0
		for {
			// item 2690002
			if i < len(path)-1 {

			} else {
				break
			}
			// item 271
			key := string(path[i]) + "-" + string(path[i+2])
			lenWay = lenWay + mapLen[key]
			priceLen = priceLen + mapLen[key]*mapPrice[key]
			// item 2690003
			i = i + 2
		}
		// item 301
		fmt.Println("\n     Длина пути |",
			math.Round(lenWay*100)/100, "км")
		fmt.Println(" Стоимость пути |",
			math.Round(priceLen*100)/100, "$")
		fmt.Println("--------------------")
	} else {

	}
	// item 280
	visit[start] = true
	// item 283
	var edge *AjlistNode = graph.node[start].next
	for {
		// item 285
		if edge != nil {

		} else {
			break
		}
		// item 287
		graph.findPath(edge.id, last,
			path+" "+
				strconv.Itoa(edge.id), visit)
		// item 288
		edge = edge.next
	}
	// item 290
	visit[start] = false
}

func getAjlistNode(id int) *AjlistNode {
	// item 57
	return &AjlistNode{id, nil}
}

func getGraph(size int) *Graph {
	// item 44
	var me *Graph = &Graph{size, make([]*Vertices, size)}
	// item 45
	me.setData()
	// item 46
	return me
}

func getVertices(data int) *Vertices {
	// item 158
	return &Vertices{data, nil, nil}
}

func main() {
	// item 4
	var graph *Graph = getGraph(7)
	// item 294
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
	// item 295
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
	// item 302
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
	// item 7
	graph.allPath(2, 5)
}