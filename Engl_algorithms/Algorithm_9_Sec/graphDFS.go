// Autogenerated with DRAKON Editor 1.31

package main
import "fmt"
/*
    Go Program for 
    Breadth first traversal in directed graph
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
    // Число вершин
    size int
    node []* Vertices
}
var k int = 0




func  (graph  *Graph) addEdge(start, last int)  {
    // item 244
    var edge *AjlistNode = 
    getAjlistNode(last)
    // item 125
    if start >= 0 && start < graph.size && 
last >= 0 && last < graph.size {
        // item 129
        if graph.node[start].next == nil {
            // item 132
            graph.node[start].next = edge
        } else {
            // item 133
            graph.node[start].last.next = edge
        }
        // item 212
        graph.node[start].last = edge
    } else {
        // item 137
        fmt.Println("\nError")
    }
}

func  (graph  *Graph) setData()  {
    // item 274
    if graph.size <= 0 {
        // item 277
        fmt.Println("\nEmpty Graph")
    } else {
        // item 2780001
        index := 0;
        for {
            // item 2780002
            if index < graph.size {
                
            } else {
                break
            }
            // item 280
            graph.node[index] =
            getVertices(index)
            // item 2780003
            index++;
        }
    }
}

func  (graph  Graph) printDFS(point int)  {
    // item 234
    if graph.size <= 0 || point < 0 || 
point >= graph.size {
        // item 237
        return
    } else {
        
    }
    // item 238
    var visit = make([]bool, graph.size)
    var path = make([] int, graph.size)
    // item 2390001
    i := 0;
    for {
        // item 2390002
        if i < graph.size {
            
        } else {
            break
        }
        // item 241
        visit[i] = false
        // item 2390003
        i++;
    }
    // item 242
    graph.dfs(visit, point, path)
    // item 282
    fmt.Println("\nDepth-first graph traversal")
    fmt.Println("path = ", path)
}

func  (graph *Graph) dfs(visit []bool, point int, path [] int)  {
    // item 218
    if visit[point] {
        // item 221
        return
    } else {
        
    }
    // item 222
    visit[point] = true
    // item 267
    path[k] = point
    k++
    // item 223
    var nodes *AjlistNode = 
    graph.node[point].next
    for {
        // item 224
        if nodes != nil {
            
        } else {
            break
        }
        // item 226
        graph.dfs(visit, nodes.id, path)
        // item 227
        nodes = nodes.next
    }
}

func  (graph Graph) printGraph()  {
    // item 151
    if graph.size > 0 {
        // item 1540001
        index := 0;
        for {
            // item 1540002
            if index < graph.size {
                
            } else {
                break
            }
            // item 156
            fmt.Print("\nAdjacency list of vertex ", 
                            index, " : ")
            // item 157
            var nodes * AjlistNode = 
            graph.node[index].next
            for {
                // item 158
                if nodes != nil {
                    
                } else {
                    break
                }
                // item 161
                fmt.Print("  ", graph.node[nodes.id].data)
                // item 162
                nodes = nodes.next
            }
            // item 1540003
            index++;
        }
    } else {
        
    }
    // item 268
    fmt.Println("")
}

func  getAjlistNode(id int) *AjlistNode {
    // item 266
    return &AjlistNode{id, nil}
}

func  getGraph(size int) * Graph {
    // item 89
    var myGraph *Graph = 
    &Graph {size,make([]*Vertices, size)}
    // item 90
    myGraph.setData()
    // item 91
    return myGraph
}

func  getVertices(data int) * Vertices {
    // item 260
    return &Vertices {data,nil,nil}
}

func  main()  {
    // item 4
    
    vert := 5 
    var graph *Graph = getGraph(vert)
    // item 5
    
           graph.addEdge(0, 2)
           graph.addEdge(0, 3)
           graph.addEdge(1, 0)
           graph.addEdge(2, 1)
           graph.addEdge(3, 4)
           graph.addEdge(4, 0)
    // item 6
    graph.printGraph()
    // item 243
    vertValue := 0
    graph.printDFS(vertValue)
}


