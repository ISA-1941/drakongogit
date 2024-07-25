// Autogenerated with DRAKON Editor 1.31

package main
import (
	"strconv"
	"fmt"
	"math"
)

/*
    Go Program for 
 Print all path between given vertices in a directed graph
*/
type AjlistNode struct {
	// Vertices node key
	id int
	next * AjlistNode
}


type Vertices struct {
	data int
	next * AjlistNode
	last * AjlistNode
}


type Graph struct {
	// Number of Vertices
	size int
	result int
	node []*Vertices
}



func  (graph  Graph) allPath(start, last int)  {
    // item 143
    if graph.size <= 0 {
        // item 146
        return
    } else {
        
    }
    // item 147
    var visit =
     make([] bool, graph.size)
    // item 1480001
    i := 0;
    for {
        // item 1480002
        if i < graph.size {
            
        } else {
            break
        }
        // item 150
        visit[i] = false
        // item 1480003
        i++;
    }
    // item 151
    fmt.Println("\nResult : ")
    // item 152
    graph.findPath(start, last, ""+ 
    strconv.Itoa(start), visit)
}

func  (graph  Graph) findPath(start, last int, path string, visit[] bool)  {
    // item 123
    f1 := start >= graph.size || last >= graph.size
    // item 124
    f2 := start < 0 || last < 0 || graph.size <= 0
    // item 120
    if f1 || f2 {
        // item 125
        return
    } else {
        
    }
    // item 126
    if start == last {
        // item 129
        fmt.Println("(", path, ")")
    } else {
        
    }
    // item 130
    visit[start] = true
    // item 131
    var edge * AjlistNode = 
    graph.node[start].next
    for {
        // item 132
        if edge != nil {
            
        } else {
            break
        }
        // item 134
        graph.findPath(edge.id, last, 
        path + " " + 
        strconv.Itoa(edge.id), visit)
        // item 135
        edge = edge.next
    }
    // item 230
    fmt.Println("findPath ",  edge)
    // item 137
    visit[start] = false
}

func  (graph  Graph) printGraph()  {
    // item 76
    if graph.size > 0 {
        // item 790001
        index := 0;
        for {
            // item 790002
            if index < graph.size {
                
            } else {
                break
            }
            // item 81
            fmt.Print("\nAdjacency list of vertex ", 
                            index, " : ")
            // item 82
            var temp * AjlistNode = 
            graph.node[index].next
            for {
                // item 83
                if temp != nil {
                    
                } else {
                    break
                }
                // item 86
                fmt.Print("  ", graph.node[temp.id].data)
                // item 87
                temp = temp.next
            }
            // item 790003
            index++;
        }
    } else {
        
    }
}

func  (graph *Graph) addEdge(start, last int)  {
    // item 13
    if start >= 0 && start < graph.size && 
last >= 0 && last < graph.size {
        // item 16
         var edge * AjlistNode = 
        getAjlistNode(last)
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

func  (graph *Graph) connect(start, last int)  {
    // item 63
     var edge * AjlistNode =
     getAjlistNode(last)
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

func  (graph *Graph) findMinimumEdge(start int, last int, visit[] bool, length int)  {
    // item 215
    f1 := start >= graph.size || last >= graph.size
    // item 216
    f2 := start < 0 || last < 0 || graph.size <= 0
    // item 195
    if f1 || f2 {
        // item 197
        return
    } else {
        
    }
    // item 198
    if visit[start] == true {
        // item 200
        return
    } else {
        
    }
    // item 201
    if start == last && length < graph.result {
        // item 203
        graph.result = length
    } else {
        
    }
    // item 204
    visit[start] = true
    // item 208
    var edge * AjlistNode = 
    graph.node[start].next
    for {
        // item 209
        if edge != nil {
            
        } else {
            break
        }
        // item 212
        graph.findMinimumEdge(edge.id, last, visit, length + 1)
        // item 213
        edge = edge.next
    }
    // item 214
    visit[start] = false
}

func  (graph *Graph) setData()  {
    // item 27
    if graph.size <= 0 {
        // item 30
        fmt.Println("\nEmpty Graph")
    } else {
        // item 310001
        index := 0;
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
            index++;
        }
    }
}

func  (graph Graph) minEdges(u,v int)  {
    // item 222
    if graph.size <= 0 {
        // item 225
        return
    } else {
        
    }
    // item 226
    var visit = make([] bool, graph.size)
    // item 227
    graph.result = math.MaxInt64
    // item 228
    graph.findMinimumEdge(u, v, visit, 0)
    // item 229
    fmt.Println("\nMinimum edges in between [", 
                    u, ",", v, "] is ", graph.result)
}

func  getAjlistNode(id int) *AjlistNode {
    // item 57
    return &AjlistNode {id,nil}
}

func  getGraph(size int) * Graph {
    // item 44
    var me *Graph = 
    &Graph {size,0, make([]*Vertices, size)}
    // item 45
    me.setData()
    // item 46
    return me
}

func  getVertices(data int) * Vertices {
    // item 158
    return &Vertices {data,nil,nil}
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
    graph.allPath(1, 4)
    graph.minEdges(2, 5)
}


